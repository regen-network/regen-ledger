# Dependency triage

How we keep on top of dependency vulnerabilities, and roughly what it costs.

Vulnerability reporting and disclosure are covered separately in
[SECURITY.md](../SECURITY.md). This document is only about routine dependency work.

## Where things live

| | |
| --- | --- |
| Dependency advisories | [Dependabot alerts](https://github.com/regen-network/regen-ledger/security/dependabot) — the source of truth |
| Source-level scanning | CodeQL default setup (Settings → Code security) |
| Monthly summary | [Security Scan workflow](workflows/security-scan.yml), posted to Slack |

Dependabot alerts are visible to Write, Maintain and Admin only. They are not public,
even though the repository is.

## Sort by blast radius, not by severity

Most bumps need no thought. Deciding which is which is where the time goes.

| Class | What is in it | How to handle it |
| --- | --- | --- |
| **Batch-merge** | `github-actions`, `/docs` npm, test-only dependencies | Merge on green CI. No review needed. |
| **Read the diff** | ordinary Go libraries | One person skims the changelog and merges on green. |
| **Careful** | anything in the `consensus` group — `cosmos-sdk`, `cometbft`, `ibc-go`, `CosmWasm`, `iavl`, `cosmos-db`, `cosmossdk.io/store`, `cosmossdk.io/x/*` | Never a routine merge. Goes into a planned upgrade, and may be state-breaking. |

`.github/dependabot.yml` groups updates to match those classes, so the group name tells
you which lane a pull request is in. Grouping applies to version updates only — security
updates stay ungrouped so a single fix can be merged on its own.

## Cadence

| When | What | Who | Time |
| --- | --- | --- | --- |
| Weekly | Merge the batch-merge and read-the-diff groups | Rotating | 30 min |
| Monthly | Triage the digest: act on critical and high, dismiss the never-fixables | Rotating | 30 min |
| Quarterly | Read release notes for the `consensus` group, looking for security fixes that were never given an advisory | **Named owner, not rotating** | 1–2 hours |

Roughly four hours a month.

## Dismissing alerts

Some advisories will never be fixed — the affected range was never capped upstream, or
the advisory describes a design limitation rather than a bug. Dismiss those in
Dependabot **with a reason**. The monthly digest excludes dismissed alerts, so the
decision holds without anyone maintaining a list.

## The Go standard library

Dependabot cannot report standard library vulnerabilities: the standard library is not
part of the dependency graph, so no alert will ever fire for one. Staying on the current
patch of our Go line is the control. The monthly digest reports when the `toolchain` pin
in `go.mod` has fallen behind.

## Checking whether a finding is reachable

Worth doing only when an alert is high or critical **and** the fix looks expensive.
Dependabot reports an advisory whether or not the vulnerable function is on a path we
execute; in one pass roughly two thirds of the advisories in the module graph were not
reachable from our code.

```sh
go install golang.org/x/vuln/cmd/govulncheck@latest

# One module at a time. Two packages in api are excluded because they do not
# compile: their generated *.pulsar.go files were never committed.
for mod in . api types x/data x/ecocredit; do
  (cd "$mod" && go list ./... |
    grep -Ev '/orm/internal/(testpb|testutil)$' |
    xargs govulncheck -scan=symbol)
done
```

`-scan=symbol` reports only advisories whose vulnerable function is reachable, and prints
the call stack. If package loading fails, the Go on your PATH may be newer than the repo
supports — pin with `GOTOOLCHAIN` to the version in `go.mod`.
