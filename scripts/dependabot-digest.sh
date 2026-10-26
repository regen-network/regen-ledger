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
  # Two expressions, not one: `+?` is not a lazy quantifier in ERE. GNU sed
  # treats it as greedy and keeps the .git suffix; BSD sed rejects it outright.
  REPO=$(git remote get-url origin 2>/dev/null |
    sed -E 's#^.*github\.com[:/]+([^/]+/[^/]+)$#\1#; s#\.git$##') || true
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

# Package paths like google.golang.org/grpc look like hostnames, so Slack
# auto-links them to the package docs, which is useless here. Making one the
# anchor text of a real link both suppresses that and points at the alert.
jq -r --arg base "${GITHUB_SERVER_URL:-https://github.com}/${REPO}/security/dependabot" '
  # The REST API reports severity as critical/high/medium/low. The Advisory
  # Database web UI says "moderate" for the same band - do not use that here.
  # Anything unrecognised is still counted and shown rather than dropped, so a
  # vocabulary change cannot silently hide alerts again.
  ["critical", "high", "medium", "low"]                     as $known
  | [.[] | select(.state == "open")]                        as $open
  | [.[] | select(.state == "dismissed" or .state == "auto_dismissed")] as $gone
  | ($open | group_by(.security_advisory.severity)
          | map({key: (.[0].security_advisory.severity // "unrated"), value: length})
          | from_entries)                                   as $counts
  | (($known | map(select($counts[.] != null)))
      + (($counts | keys) - $known | sort))                 as $order
  | ($order | map("\($counts[.]) \(.)") | join(", "))       as $headline

  # One line per advisory, not per alert: the same advisory raises a separate
  # alert for every manifest it affects, so five go.mod files mean five alerts
  # and one bump. Grouping keeps the message short enough not to be split.
  | ([$open[] | select(.security_advisory.severity == "critical")]
      | group_by(.security_advisory.ghsa_id)
      | map(
          (.[0]) as $a
          | "• <\($base)/\($a.number)|\($a.security_vulnerability.package.name)> "
            + "\($a.security_advisory.ghsa_id) — fix "
            + ($a.security_vulnerability.first_patched_version.identifier
               // "not yet available")
            + (if length > 1 then " (\(length) alerts)" else "" end)
        )
      | sort)                                               as $criticals
  | if ($open | length) == 0 then
      "No open Dependabot alerts."
    else
      [ "*\($headline)* open Dependabot alert(s)." ]
      + (if ($criticals | length) > 0 then
          ["Critical, by advisory:"] + $criticals
         else [] end)
      + (if ($gone | length) > 0 then
          ["\($gone | length) dismissed and excluded."]
         else [] end)
      | join("\n")
    end
' < "$alerts"

# Dependabot cannot report Go standard library vulnerabilities: the standard
# library is not part of the dependency graph, so no alert will ever fire for
# it. Staying on the current patch of our Go line is the control, so the digest
# says when the pin has fallen behind. Network failure here is not fatal - the
# alert summary above still stands.
repo_root="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
pinned=$(grep -E "^toolchain " "${repo_root}/go.mod" 2>/dev/null | awk '{print $2}')

if [[ -n "$pinned" ]]; then
  minor=${pinned#go}
  minor=${minor%.*}
  latest=$(curl -sS --max-time 15 "https://go.dev/dl/?mode=json&include=all" 2>/dev/null |
    jq -r --arg m "go${minor}." '
      [.[] | select(.stable) | .version | select(startswith($m))]
      | sort_by(ltrimstr("go") | split(".") | map(tonumber? // 0))
      | last // empty' 2>/dev/null) || latest=""

  if [[ -n "$latest" && "$latest" != "$pinned" ]]; then
    echo "Go toolchain pinned at ${pinned}; ${latest} is released. Standard library"\
         "fixes ship in Go releases and raise no Dependabot alert."
  fi
fi
