// Command govulncheck-severity rates, ranks and renders govulncheck SARIF output.
//
// govulncheck reports no severity: the Go vulnerability database leaves the OSV
// severity field empty on every entry. Its entries do carry GHSA aliases, so
// severity is recoverable by resolving GO-ID -> GHSA and looking the advisory
// up. The same OSV lookup yields fixed versions, which SARIF omits.
//
// Severity is taken from this repository's Dependabot alerts where possible:
// one paginated request covers every advisory, it is the same data the team
// already triages, and it carries dismissal state, so anything dismissed there
// stops being counted here. Where an advisory has no Dependabot alert - or the
// alert list cannot be read - it falls back to the public GitHub Advisory
// Database, one request per advisory.
//
// One scan, three outputs:
//
//	-out-text   a severity-ranked report, most serious first
//	-out-json   a severity map, consumed by -aggregate for the alert
//
// Findings are not uploaded to code scanning: third-party advisories are
// already tracked as Dependabot alerts, and duplicating them into a second
// alert inbox costs more than it adds. What Dependabot cannot report is the
// standard library and whether a vulnerable symbol is actually called, so the
// reports call both out explicitly.
//
// Go standard library findings have no GHSA alias and stay unrated rather than
// being given an invented score. They are almost always closed by a single Go
// toolchain bump, which the reports state separately.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	goDBURL     = "https://vuln.go.dev/ID/%s.json"
	advisoryURL = "https://api.github.com/advisories/%s"
)

// severityOrder is both the report ordering and the aggregation ordering.
var severityOrder = []string{"critical", "high", "medium", "low", "unrated"}

// bucket maps a CVSS score onto GitHub code scanning's severity bands.
func bucket(score float64) string {
	switch {
	case score >= 9.0:
		return "critical"
	case score >= 7.0:
		return "high"
	case score >= 4.0:
		return "medium"
	default:
		return "low"
	}
}

// finding is the per-advisory rating written to the severity map. The json
// names are load-bearing: -aggregate reads them back.
type finding struct {
	Severity string   `json:"severity"`
	CVSS     *float64 `json:"cvss"`
	GHSA     string   `json:"ghsa,omitempty"`
	Package  string   `json:"package,omitempty"`
	Fixed    []string `json:"fixed,omitempty"`
	// Dismissed mirrors a Dependabot alert that someone has triaged away, so a
	// decision made once there is not re-raised here every month.
	Dismissed bool `json:"dismissed,omitempty"`
	// Source records where the rating came from: dependabot, advisory, or none.
	Source string `json:"source,omitempty"`
}

type severityMap struct {
	Module   string             `json:"module"`
	Findings map[string]finding `json:"findings"`
}

type osvEntry struct {
	Aliases  []string `json:"aliases"`
	Affected []struct {
		Package struct {
			Name string `json:"name"`
		} `json:"package"`
		Ranges []struct {
			Events []map[string]string `json:"events"`
		} `json:"ranges"`
	} `json:"affected"`
}

// dependabotAlert is the subset of GET /repos/{owner}/{repo}/dependabot/alerts
// this command uses. Reading it needs the `vulnerability-alerts: read`
// permission, which is separate from `security-events`.
type dependabotAlert struct {
	State            string `json:"state"`
	SecurityAdvisory struct {
		GHSAID         string `json:"ghsa_id"`
		Severity       string `json:"severity"`
		CVSSSeverities map[string]struct {
			Score float64 `json:"score"`
		} `json:"cvss_severities"`
	} `json:"security_advisory"`
	SecurityVulnerability struct {
		Package struct {
			Name string `json:"name"`
		} `json:"package"`
	} `json:"security_vulnerability"`
}

type ghAdvisory struct {
	Severity       string `json:"severity"`
	CVSSSeverities map[string]struct {
		Score float64 `json:"score"`
	} `json:"cvss_severities"`
	Vulnerabilities []struct {
		Package struct {
			Name string `json:"name"`
		} `json:"package"`
	} `json:"vulnerabilities"`
}

type rater struct {
	client *http.Client
	token  string
	cache  map[string]finding
	// alerts is keyed by GHSA id; empty when Dependabot alerts are unreadable.
	alerts map[string]dependabotAlert
}

func newRater() *rater {
	return &rater{
		client: &http.Client{Timeout: 30 * time.Second},
		token:  os.Getenv("GITHUB_TOKEN"),
		cache:  map[string]finding{},
		alerts: map[string]dependabotAlert{},
	}
}

// loadDependabotAlerts indexes this repository's Dependabot alerts by GHSA id.
// Failure is not fatal: severity then falls back to per-advisory lookups, which
// work without any repository permission at all.
func (r *rater) loadDependabotAlerts(repo string) {
	if repo == "" || r.token == "" {
		fmt.Fprintln(os.Stderr, "dependabot: no repository or token, using advisory lookups")
		return
	}
	for page := 1; ; page++ {
		var batch []dependabotAlert
		url := fmt.Sprintf("https://api.github.com/repos/%s/dependabot/alerts?per_page=100&page=%d", repo, page)
		if err := r.get(url, &batch); err != nil {
			fmt.Fprintf(os.Stderr, "dependabot: %v; using advisory lookups instead\n", err)
			r.alerts = map[string]dependabotAlert{}
			return
		}
		for _, alert := range batch {
			if id := alert.SecurityAdvisory.GHSAID; id != "" {
				// keep the first, which is the most recent alert for the advisory
				if _, seen := r.alerts[id]; !seen {
					r.alerts[id] = alert
				}
			}
		}
		if len(batch) < 100 {
			break
		}
	}
	fmt.Fprintf(os.Stderr, "dependabot: indexed %d alert(s) for severity\n", len(r.alerts))
}

func (r *rater) get(url string, out any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "regen-ledger-security-scan")
	if r.token != "" && strings.Contains(url, "api.github.com") {
		req.Header.Set("Authorization", "Bearer "+r.token)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck // read-only request

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s: %s", url, resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

// rate resolves one GO-ID to a severity, CVSS score, package and fixed versions.
// Lookup failures are reported but never fatal: a finding stays unrated and the
// report says so, so throttling cannot break the scan.
func (r *rater) rate(goID string) finding {
	if cached, ok := r.cache[goID]; ok {
		return cached
	}

	result := finding{Severity: "unrated"}

	var entry osvEntry
	if err := r.get(fmt.Sprintf(goDBURL, goID), &entry); err != nil {
		fmt.Fprintf(os.Stderr, "%s: vulnerability lookup failed (%v)\n", goID, err)
		r.cache[goID] = result
		return result
	}

	// Fixed versions come from the OSV record; SARIF does not carry them.
	seen := map[string]bool{}
	for _, affected := range entry.Affected {
		name := affected.Package.Name
		if result.Package == "" {
			result.Package = name
		}
		for _, rng := range affected.Ranges {
			for _, event := range rng.Events {
				if fixed := event["fixed"]; fixed != "" {
					if version := name + "@" + fixed; !seen[version] {
						seen[version] = true
						result.Fixed = append(result.Fixed, version)
					}
				}
			}
		}
	}
	sort.Strings(result.Fixed)

	var ghsa string
	for _, alias := range entry.Aliases {
		if strings.HasPrefix(alias, "GHSA-") {
			ghsa = alias
			break
		}
	}
	if ghsa == "" {
		r.cache[goID] = result
		return result
	}
	result.GHSA = ghsa

	// Prefer the repository's own Dependabot alert: same advisory data, one
	// request for all of them, and it knows what has already been dismissed.
	if alert, ok := r.alerts[ghsa]; ok {
		result.Source = "dependabot"
		result.Dismissed = alert.State == "dismissed" || alert.State == "auto_dismissed"
		for _, key := range []string{"cvss_v4", "cvss_v3"} {
			if value, ok := alert.SecurityAdvisory.CVSSSeverities[key]; ok && value.Score > 0 {
				score := value.Score
				result.CVSS = &score
				break
			}
		}
		switch {
		case result.CVSS != nil:
			result.Severity = bucket(*result.CVSS)
		case alert.SecurityAdvisory.Severity != "":
			result.Severity = alert.SecurityAdvisory.Severity
		}
		if result.Package == "" {
			result.Package = alert.SecurityVulnerability.Package.Name
		}
		r.cache[goID] = result
		return result
	}

	var advisory ghAdvisory
	if err := r.get(fmt.Sprintf(advisoryURL, ghsa), &advisory); err != nil {
		fmt.Fprintf(os.Stderr, "%s: advisory lookup failed (%v)\n", goID, err)
		r.cache[goID] = result
		return result
	}

	result.Source = "advisory"
	for _, key := range []string{"cvss_v4", "cvss_v3"} {
		if value, ok := advisory.CVSSSeverities[key]; ok && value.Score > 0 {
			score := value.Score
			result.CVSS = &score
			break
		}
	}
	// Prefer the band implied by the score; fall back to GitHub's own label.
	switch {
	case result.CVSS != nil:
		result.Severity = bucket(*result.CVSS)
	case advisory.Severity != "":
		result.Severity = advisory.Severity
	}

	r.cache[goID] = result
	return result
}

// Generic JSON navigation. The SARIF document is held as decoded maps rather
// than typed structs so that fields this command does not model - code flows,
// call stacks, artifacts - survive the round-trip untouched.

func asMap(v any) map[string]any {
	m, _ := v.(map[string]any)
	return m
}

func asSlice(v any) []any {
	s, _ := v.([]any)
	return s
}

func asString(v any) string {
	s, _ := v.(string)
	return s
}

func dig(v any, keys ...string) any {
	for _, key := range keys {
		m := asMap(v)
		if m == nil {
			return nil
		}
		v = m[key]
	}
	return v
}

// traces flattens a result's first call stack into readable file:line entries.
func traces(result map[string]any, limit int) []string {
	var out []string
	stacks := asSlice(result["stacks"])
	if len(stacks) == 0 {
		return nil
	}
	for _, frame := range asSlice(dig(stacks[0], "frames")) {
		if len(out) >= limit {
			break
		}
		location := dig(frame, "location")
		uri := asString(dig(location, "physicalLocation", "artifactLocation", "uri"))
		if uri == "" {
			uri = "?"
		}
		line := "?"
		if raw := dig(location, "physicalLocation", "region", "startLine"); raw != nil {
			line = fmt.Sprint(raw)
		}
		symbol := asString(dig(location, "message", "text"))
		out = append(out, strings.TrimSpace(fmt.Sprintf("%s:%s %s", uri, line, symbol)))
	}
	return out
}

func countsLine(counts map[string]int, includeUnrated bool) string {
	var parts []string
	for _, severity := range severityOrder {
		if severity == "unrated" && !includeUnrated {
			continue
		}
		if counts[severity] > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", counts[severity], severity))
		}
	}
	return strings.Join(parts, ", ")
}

// render writes a severity-ranked report, worst first, with stdlib grouped last.
func render(module string, rules map[string]map[string]any, results []any, ratings map[string]finding) string {
	byRule := map[string][]map[string]any{}
	reachable := 0
	for _, raw := range results {
		result := asMap(raw)
		if asString(result["level"]) != "error" {
			continue
		}
		reachable++
		id := asString(result["ruleId"])
		byRule[id] = append(byRule[id], result)
	}

	grouped := map[string][]string{}
	counts := map[string]int{}
	dismissed := 0
	for id := range byRule {
		if ratings[id].Dismissed {
			dismissed++
			continue
		}
		severity := "unrated"
		if rating, ok := ratings[id]; ok && rating.Severity != "" {
			severity = rating.Severity
		}
		grouped[severity] = append(grouped[severity], id)
		counts[severity]++
	}

	var b strings.Builder
	fmt.Fprintf(&b, "govulncheck report — %s\n", module)
	fmt.Fprintf(&b, "%d reachable finding(s) of %d total.\n\n", reachable, len(results))
	if line := countsLine(counts, true); line != "" {
		fmt.Fprintf(&b, "By severity: %s\n", line)
	}
	stdlib := 0
	for id := range byRule {
		if ratings[id].Package == "stdlib" && !ratings[id].Dismissed {
			stdlib++
		}
	}
	fmt.Fprintf(&b, "Standard library: %d, third-party: %d. Third-party advisories are "+
		"also tracked as Dependabot alerts; the standard library is not.\n",
		stdlib, len(byRule)-stdlib-dismissed)
	if dismissed > 0 {
		fmt.Fprintf(&b, "%d finding(s) dismissed in Dependabot are excluded.\n", dismissed)
	}
	b.WriteString("\n")

	for _, severity := range severityOrder {
		ids := grouped[severity]
		if len(ids) == 0 {
			continue
		}
		heading := strings.ToUpper(severity)
		if severity == "unrated" {
			heading += " (mostly Go standard library — closed by a toolchain bump)"
		}
		fmt.Fprintf(&b, "## %s\n\n", heading)

		// worst CVSS first, then by id so output is stable
		sort.Slice(ids, func(i, j int) bool {
			left, right := ratings[ids[i]], ratings[ids[j]]
			li, ri := 0.0, 0.0
			if left.CVSS != nil {
				li = *left.CVSS
			}
			if right.CVSS != nil {
				ri = *right.CVSS
			}
			if li != ri {
				return li > ri
			}
			return ids[i] < ids[j]
		})

		for _, id := range ids {
			rating := ratings[id]
			score := "no CVSS"
			if rating.CVSS != nil {
				score = "CVSS " + strconv.FormatFloat(*rating.CVSS, 'f', -1, 64)
			}
			fmt.Fprintf(&b, "%s  [%s]  %s\n", id, score, rating.Package)
			if title := asString(dig(rules[id], "fullDescription", "text")); title != "" {
				fmt.Fprintf(&b, "  %s\n", title)
			}
			fixed := "no fixed version published"
			if len(rating.Fixed) > 0 {
				fixed = strings.Join(rating.Fixed, ", ")
			}
			fmt.Fprintf(&b, "  Fixed in: %s\n", fixed)
			for _, trace := range traces(byRule[id][0], 3) {
				fmt.Fprintf(&b, "  via %s\n", trace)
			}
			if uri := asString(dig(rules[id], "helpUri")); uri != "" {
				fmt.Fprintf(&b, "  %s\n", uri)
			}
			b.WriteString("\n")
		}
	}
	return strings.TrimRight(b.String(), "\n") + "\n"
}

// aggregate combines per-module severity maps into one weighted headline.
func aggregate(dir string) (string, bool, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", false, err
	}

	counts := map[string]int{}
	criticals := map[string]finding{}
	var uncovered []string
	// Tracked separately because these are the two things Dependabot alerts
	// cannot tell you: the standard library is not in the dependency graph, and
	// nothing there reports whether a vulnerable symbol is actually called.
	stdlib, thirdParty, dismissed := 0, 0, 0
	thirdPartyPkgs := map[string]bool{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		module := strings.TrimPrefix(entry.Name(), "govulncheck-")
		path := filepath.Join(dir, entry.Name(), "govulncheck-severity.json")

		// No severity map means the scan never produced one, so that module was
		// not covered. Treating it as clean would hide the gap.
		raw, err := os.ReadFile(path)
		if err != nil {
			uncovered = append(uncovered, module)
			continue
		}
		var parsed severityMap
		if err := json.Unmarshal(raw, &parsed); err != nil {
			uncovered = append(uncovered, module)
			continue
		}
		for id, rating := range parsed.Findings {
			// Dismissed in Dependabot means someone has already triaged it away;
			// re-raising it every month is what makes a monthly alert ignorable.
			if rating.Dismissed {
				dismissed++
				continue
			}
			severity := rating.Severity
			if severity == "" {
				severity = "unrated"
			}
			counts[severity]++
			if rating.Package == "stdlib" {
				stdlib++
			} else {
				thirdParty++
				if rating.Package != "" {
					thirdPartyPkgs[rating.Package] = true
				}
			}
			if severity == "critical" {
				criticals[id] = rating
			}
		}
	}

	var lines []string
	rated := countsLine(counts, false)
	if rated != "" {
		lines = append(lines, fmt.Sprintf("*%s* need engineering review.", rated))
	}

	ids := make([]string, 0, len(criticals))
	for id := range criticals {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		rating := criticals[id]
		pkg := rating.Package
		if pkg == "" {
			pkg = "see report"
		}
		fixed := "no fix published"
		if len(rating.Fixed) > 0 {
			fixed = strings.Join(rating.Fixed, ", ")
		}
		score := "unknown"
		if rating.CVSS != nil {
			score = strconv.FormatFloat(*rating.CVSS, 'f', -1, 64)
		}
		lines = append(lines, fmt.Sprintf("Critical: %s (%s) — CVSS %s, fixed in %s", id, pkg, score, fixed))
	}

	if stdlib > 0 {
		lines = append(lines, fmt.Sprintf(
			"%d of %d are Go standard library — one toolchain bump closes these, and "+
				"Dependabot cannot see them.", stdlib, stdlib+thirdParty))
	}
	if thirdParty > 0 {
		lines = append(lines, fmt.Sprintf(
			"%d are third-party across %d module(s), also tracked in Dependabot alerts.",
			thirdParty, len(thirdPartyPkgs)))
	}
	if dismissed > 0 {
		lines = append(lines, fmt.Sprintf(
			"%d dismissed in Dependabot and excluded from the counts above.", dismissed))
	}
	if len(uncovered) > 0 {
		sort.Strings(uncovered)
		lines = append(lines, "Not covered by this scan: "+strings.Join(uncovered, ", "))
	}

	body := strings.Join(lines, "\n")
	found := rated != "" || counts["unrated"] > 0 || len(uncovered) > 0

	if out := os.Getenv("GITHUB_OUTPUT"); out != "" {
		file, err := os.OpenFile(out, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {
			return body, found, err
		}
		defer file.Close() //nolint:errcheck // best-effort append
		if _, err := fmt.Fprintf(file, "found=%t\nheadline=%s\nbody<<REPORT_EOF\n%s\nREPORT_EOF\n",
			found, rated, body); err != nil {
			return body, found, err
		}
	}
	return body, found, nil
}

func writeJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644) //nolint:gosec // reports are not secret
}

func main() {
	sarifPath := flag.String("sarif", "", "SARIF file to read")
	outJSON := flag.String("out-json", "", "write the severity map here")
	outText := flag.String("out-text", "", "write the ranked report here")
	module := flag.String("module", "module", "label used in reports")
	aggregateDir := flag.String("aggregate", "", "combine per-module maps in this directory")
	flag.Parse()

	log.SetFlags(0)

	if *aggregateDir != "" {
		body, _, err := aggregate(*aggregateDir)
		if err != nil {
			log.Fatalf("aggregate: %v", err)
		}
		if body == "" {
			body = "No findings in any module."
		}
		fmt.Println(body)
		return
	}

	info, err := os.Stat(*sarifPath)
	if *sarifPath == "" || err != nil || info.Size() == 0 {
		log.Fatalf("no SARIF at %q, nothing to rate", *sarifPath)
	}

	file, err := os.Open(*sarifPath)
	if err != nil {
		log.Fatalf("open SARIF: %v", err)
	}
	decoder := json.NewDecoder(file)
	// Keep numbers verbatim so line numbers and scores round-trip unchanged.
	decoder.UseNumber()
	var doc map[string]any
	if err := decoder.Decode(&doc); err != nil {
		file.Close() //nolint:errcheck // already failing
		log.Fatalf("decode SARIF: %v", err)
	}
	if err := file.Close(); err != nil {
		log.Fatalf("close SARIF: %v", err)
	}

	r := newRater()
	r.loadDependabotAlerts(os.Getenv("GITHUB_REPOSITORY"))
	ratings := map[string]finding{}
	rules := map[string]map[string]any{}
	var results []any

	for _, rawRun := range asSlice(doc["runs"]) {
		run := asMap(rawRun)
		runResults := asSlice(run["results"])
		results = append(results, runResults...)

		reachable := map[string]bool{}
		for _, rawResult := range runResults {
			result := asMap(rawResult)
			if asString(result["level"]) == "error" {
				reachable[asString(result["ruleId"])] = true
			}
		}

		for _, rawRule := range asSlice(dig(run, "tool", "driver", "rules")) {
			rule := asMap(rawRule)
			id := asString(rule["id"])
			if id == "" {
				continue
			}
			rules[id] = rule

			rating := r.rate(id)
			if reachable[id] {
				ratings[id] = rating
			}

		}
	}

	if *outJSON != "" {
		if err := writeJSON(*outJSON, severityMap{Module: *module, Findings: ratings}); err != nil {
			log.Fatalf("write severity map: %v", err)
		}
	}
	if *outText != "" {
		report := render(*module, rules, results, ratings)
		if err := os.WriteFile(*outText, []byte(report), 0o644); err != nil { //nolint:gosec // reports are not secret
			log.Fatalf("write report: %v", err)
		}
	}

	counts := map[string]int{}
	for _, rating := range ratings {
		counts[rating.Severity]++
	}
	summary := countsLine(counts, true)
	if summary == "" {
		summary = "no reachable findings"
	}
	fmt.Printf("%s: %s\n", *module, summary)
}
