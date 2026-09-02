#!/usr/bin/env bash

# Scans every go module in the repo against the Go vulnerability database
# (https://vuln.go.dev) using govulncheck.
#
# Exits 1 if any module reports a vulnerability, so CI can fail the job.
#
# Environment variables:
#   SCAN_MODE  "symbol" (default) reports only vulnerabilities reachable from
#              this code; "package" reports advisories in imported packages;
#              "module" reports every advisory in the module graph without
#              reachability analysis, which is much noisier.
#   FORMAT     "ranked" (default) scans as SARIF and renders a report ordered
#              by severity, matching what CI produces; also "text", "json" or
#              "sarif" for raw govulncheck output. Ranking uses
#              scripts/govulncheck-severity and needs the go toolchain.
#   GITHUB_TOKEN
#              raises the advisory-lookup rate limit from 60/hour. Without it,
#              throttled findings degrade to "unrated" rather than failing.
#   OUTPUT_DIR when set, per-module reports are written here instead of stdout.
#   EXCLUDE_PACKAGES_REGEX
#              grep -E regex for package import paths excluded from symbol and
#              package scans.

set -uo pipefail

REPO_ROOT="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )"

SCAN_MODE="${SCAN_MODE:-symbol}"
FORMAT="${FORMAT:-ranked}"
OUTPUT_DIR="${OUTPUT_DIR:-}"
EXCLUDE_PACKAGES_REGEX="${EXCLUDE_PACKAGES_REGEX:-/orm/internal/(testpb|testutil)$}"

if [[ "$SCAN_MODE" == "source" ]]; then
  echo "SCAN_MODE=source is deprecated; using SCAN_MODE=symbol" >&2
  SCAN_MODE="symbol"
fi

case "$SCAN_MODE" in
  symbol|package|module) ;;
  *)
    echo "invalid SCAN_MODE ${SCAN_MODE}: expected symbol, package, or module" >&2
    exit 2
    ;;
esac

if ! command -v govulncheck > /dev/null; then
  echo "govulncheck not found, install it with 'make vulncheck-install'" >&2
  exit 1
fi

if [[ -n "$OUTPUT_DIR" ]]; then
  mkdir -p "$OUTPUT_DIR"
fi

# "ranked" is not a govulncheck format: it scans as SARIF and renders the
# result through scripts/govulncheck-severity, which recovers severity from
# the GHSA aliases. Module mode produces no reachability data to rank, so it
# falls back to text.
RANK=""
govulncheck_format="$FORMAT"
if [[ "$FORMAT" == "ranked" ]]; then
  if [[ "$SCAN_MODE" == "module" ]]; then
    echo "FORMAT=ranked needs reachability data; using text for SCAN_MODE=module" >&2
    govulncheck_format="text"
  elif ! command -v go > /dev/null; then
    echo "FORMAT=ranked needs the go toolchain; using text" >&2
    govulncheck_format="text"
  else
    RANK=1
    govulncheck_format="sarif"
  fi
fi

# Ranked rendering needs the SARIF in a file. With no OUTPUT_DIR the caller
# only wants terminal output, so stage it in a temp dir that is cleaned up.
RANK_TMPDIR=""
RANK_BIN=""
if [[ -n "$RANK" ]]; then
  RANK_TMPDIR=$(mktemp -d)
  trap 'rm -rf "$RANK_TMPDIR"' EXIT
  RANK_BIN="${RANK_TMPDIR}/govulncheck-severity"
  if ! (cd "$REPO_ROOT" && go build -o "$RANK_BIN" ./scripts/govulncheck-severity); then
    echo "could not build the severity tool; using text" >&2
    RANK=""
    RANK_BIN=""
    govulncheck_format="text"
    scan_args=("-scan=${SCAN_MODE}" "-format=text")
  fi
  if [[ -n "$RANK" && -z "$OUTPUT_DIR" ]]; then
    OUTPUT_DIR="$RANK_TMPDIR"
  fi
fi

scan_args=("-scan=${SCAN_MODE}" "-format=${govulncheck_format}")

failed=()

for modfile in $(find "${REPO_ROOT}" -name go.mod -not -path "*/node_modules/*" | sort); do
  dir=$(dirname "$modfile")
  module=$(grep "^module" "$modfile" | awk '{print $2}')
  # path relative to the repo root, used for report file names
  rel="${dir#"${REPO_ROOT}"}"
  rel="${rel#/}"
  rel="${rel:-root}"

  echo "==> scanning ${module} [$(date -u +%Y-%m-%dT%H:%M:%SZ)]"

  # In module mode govulncheck takes no package pattern and loads the working
  # directory, which fails for modules whose go files all live in subdirectories
  # (api). Scanning from any directory in the module covers the same go.mod.
  scan_dir="$dir"
  if [[ "$SCAN_MODE" == "module" ]] && ! compgen -G "${dir}/*.go" > /dev/null; then
    scan_dir=$(dirname "$(find "$dir" -name '*.go' -not -path '*/node_modules/*' | sort | head -1)")
  fi

  packages_file=""
  list_err=""
  packages=()
  if [[ "$SCAN_MODE" != "module" ]]; then
    packages_file=$(mktemp)
    list_err=$(mktemp)
    if ! (cd "$scan_dir" && go list ./...) > "$packages_file" 2> "$list_err"; then
      status=1
      if [[ -n "$OUTPUT_DIR" ]]; then
        report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${govulncheck_format}"
        cat "$list_err" > "$report"
        cat "$report"
      else
        cat "$list_err" >&2
      fi
      failed+=("${module} (exit ${status})")
      rm -f "$packages_file" "$list_err"
      continue
    fi
    rm -f "$list_err"

    if [[ -n "$EXCLUDE_PACKAGES_REGEX" ]]; then
      grep -Ev "$EXCLUDE_PACKAGES_REGEX" "$packages_file" > "${packages_file}.filtered" || true
      mv "${packages_file}.filtered" "$packages_file"
    fi

    if [[ ! -s "$packages_file" ]]; then
      status=1
      if [[ -n "$OUTPUT_DIR" ]]; then
        report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${govulncheck_format}"
        echo "no packages left to scan after applying EXCLUDE_PACKAGES_REGEX" > "$report"
        cat "$report"
      else
        echo "no packages left to scan after applying EXCLUDE_PACKAGES_REGEX" >&2
      fi
      failed+=("${module} (exit ${status})")
      rm -f "$packages_file"
      continue
    fi

    while IFS= read -r package; do
      packages+=("$package")
    done < "$packages_file"
  fi

  if [[ -n "$OUTPUT_DIR" ]]; then
    report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${govulncheck_format}"
    if [[ "$SCAN_MODE" == "module" ]]; then
      (cd "$scan_dir" && govulncheck "${scan_args[@]}") > "$report" 2>&1
    else
      (cd "$scan_dir" && govulncheck "${scan_args[@]}" "${packages[@]}") > "$report" 2>&1
    fi
    status=$?
    # Under FORMAT=ranked $report holds SARIF, which is unreadable; the ranked
    # rendering below is printed instead.
    if [[ -z "$RANK" ]]; then
      cat "$report"
    fi
  else
    if [[ "$SCAN_MODE" == "module" ]]; then
      (cd "$scan_dir" && govulncheck "${scan_args[@]}")
    else
      (cd "$scan_dir" && govulncheck "${scan_args[@]}" "${packages[@]}")
    fi
    status=$?
  fi
  rm -f "$packages_file"

  if [[ -n "$RANK" && -s "$report" ]]; then
    # $report holds SARIF at this point; render it in place, worst first.
    ranked="${OUTPUT_DIR}/govulncheck-${rel//\//-}.txt"
    if "$RANK_BIN" \
        -sarif "$report" -out-text "$ranked" \
        -out-json "${OUTPUT_DIR}/govulncheck-${rel//\//-}-severity.json" \
        -module "$rel"; then
      cat "$ranked"
    fi
  fi

  # govulncheck exits 3 when it finds vulnerabilities, 0 when clean and
  # 1 on an internal error. Treat anything non-zero as a failure but keep
  # scanning the remaining modules.
  if [[ $status -ne 0 ]]; then
    failed+=("${module} (exit ${status})")
  fi
done

if [[ ${#failed[@]} -gt 0 ]]; then
  echo
  echo "vulnerabilities or errors reported in ${#failed[@]} module(s):"
  for f in "${failed[@]}"; do
    echo "  - ${f}"
  done
  exit 1
fi

echo
echo "no vulnerabilities found in any module"
