#!/usr/bin/env bash

# Summarises this repository's open Dependabot alerts as a short digest,
# grouped by severity, with anything critical named.
#
# Dependabot already alerts per advisory as they are published. This exists
# only to give a periodic "where do we stand" summary that can be posted to a
# channel, rather than to detect anything Dependabot does not already know.
#
# Environment variables:
#   GITHUB_TOKEN  required. Needs the `vulnerability-alerts: read` permission
#                 in Actions, or the security_events scope for a PAT.
#   GITHUB_REPOSITORY
#                 owner/name. Derived from the origin remote when unset.

set -uo pipefail

TOKEN="${GITHUB_TOKEN:-${GH_TOKEN:-}}"
REPO="${GITHUB_REPOSITORY:-}"

if [[ -z "$REPO" ]]; then
  REPO=$(git remote get-url origin 2>/dev/null |
    sed -E 's#^.*github\.com[:/]+([^/]+/[^/]+?)(\.git)?$#\1#') || true
fi

if [[ -z "$TOKEN" || -z "$REPO" ]]; then
  echo "GITHUB_TOKEN and GITHUB_REPOSITORY (or an origin remote) are required" >&2
  exit 2
fi

alerts=$(mktemp)
page=$(mktemp)
headers=$(mktemp)
trap 'rm -f "$alerts" "$page" "$headers"' EXIT
echo '[]' > "$alerts"

# This endpoint paginates by cursor, not by page number: passing `page` is
# rejected with a 400. The next cursor arrives in the Link header, so follow
# that until there is no rel="next".
#
# Alerts are fetched in every state so dismissals can be counted rather than
# quietly missing from the totals.
url="https://api.github.com/repos/${REPO}/dependabot/alerts?per_page=100"

while [[ -n "$url" ]]; do
  if ! curl -sSf -D "$headers" \
    -H "Authorization: Bearer ${TOKEN}" \
    -H "Accept: application/vnd.github+json" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    "$url" > "$page"; then
    echo "could not read Dependabot alerts for ${REPO}" >&2
    echo "check that Dependabot alerts are enabled and the token grants vulnerability-alerts: read" >&2
    exit 1
  fi

  if ! jq -e 'type == "array"' > /dev/null < "$page"; then
    echo "unexpected response: $(jq -r '.message // "not an array"' < "$page")" >&2
    exit 1
  fi

  jq -s 'add' "$alerts" "$page" > "${alerts}.merged" && mv "${alerts}.merged" "$alerts"

  # One rel per line so a Link header carrying prev/first/last cannot be
  # mistaken for next.
  url=$(tr ',' '\n' < "$headers" |
    sed -n 's/.*<\([^>]*\)>[[:space:]]*;[[:space:]]*rel="next".*/\1/p' |
    head -1)
done

jq -r '
  # Dependabot severities, worst first. Note "moderate", not "medium".
  ["critical", "high", "moderate", "low"] as $order
  | [.[] | select(.state == "open")]                        as $open
  | [.[] | select(.state == "dismissed" or .state == "auto_dismissed")] as $gone
  | ($open | group_by(.security_advisory.severity)
          | map({key: .[0].security_advisory.severity, value: length})
          | from_entries)                                   as $counts
  | ($order | map(select($counts[.] != null)
          | "\($counts[.]) \(.)") | join(", "))             as $headline
  | ([$open[] | select(.security_advisory.severity == "critical")
      | "Critical: \(.security_vulnerability.package.name) "
        + "(\(.security_advisory.ghsa_id)) — fix "
        + ((.security_vulnerability.first_patched_version.identifier // "not yet available"))
     ])                                                     as $criticals
  | if ($open | length) == 0 then
      "No open Dependabot alerts."
    else
      [ "*\($headline)* open Dependabot alert(s)." ]
      + $criticals
      + (if ($gone | length) > 0 then
          ["\($gone | length) dismissed and excluded."]
         else [] end)
      | join("\n")
    end
' < "$alerts"
