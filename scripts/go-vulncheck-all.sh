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
#   FORMAT     govulncheck output format: "text" (default), "json" or "sarif".
#   OUTPUT_DIR when set, per-module reports are written here instead of stdout.
#   EXCLUDE_PACKAGES_REGEX
#              grep -E regex for package import paths excluded from symbol and
#              package scans.

set -uo pipefail

REPO_ROOT="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )"

SCAN_MODE="${SCAN_MODE:-symbol}"
FORMAT="${FORMAT:-text}"
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

scan_args=("-scan=${SCAN_MODE}" "-format=${FORMAT}")

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
        report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${FORMAT}"
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
        report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${FORMAT}"
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
    report="${OUTPUT_DIR}/govulncheck-${rel//\//-}.${FORMAT}"
    if [[ "$SCAN_MODE" == "module" ]]; then
      (cd "$scan_dir" && govulncheck "${scan_args[@]}") > "$report" 2>&1
    else
      (cd "$scan_dir" && govulncheck "${scan_args[@]}" "${packages[@]}") > "$report" 2>&1
    fi
    status=$?
    cat "$report"
  else
    if [[ "$SCAN_MODE" == "module" ]]; then
      (cd "$scan_dir" && govulncheck "${scan_args[@]}")
    else
      (cd "$scan_dir" && govulncheck "${scan_args[@]}" "${packages[@]}")
    fi
    status=$?
  fi
  rm -f "$packages_file"

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
