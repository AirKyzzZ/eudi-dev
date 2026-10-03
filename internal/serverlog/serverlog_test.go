// Copyright 2026 Dominik Schlosser
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serverlog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func records(t *testing.T, out string) []map[string]any {
	t.Helper()
	var recs []map[string]any
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		if line == "" {
			continue
		}
		var rec map[string]any
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			t.Fatalf("not a JSON record: %q", line)
		}
		recs = append(recs, rec)
	}
	return recs
}

func TestWriterRecords(t *testing.T) {
	cases := []struct {
		line      string
		level     string
		msg       string
		component string
	}{
		{"[VCI] WARNING: HAIP violation: x\n", "WARN", "HAIP violation: x", "VCI"},
		{"[Demo issuer] issuing a code\n", "INFO", "issuing a code", "Demo issuer"},
		{"  ERROR: saving log entry: disk full\n", "ERROR", "saving log entry: disk full", ""},
		{"  Warning:     base URL mismatch\n", "WARN", "base URL mismatch", ""},
		{"warning: rehydrating credential 1\n", "WARN", "rehydrating credential 1", ""},
		{"  Consent:       approved\n", "INFO", "Consent: approved", ""},
		{"  [mso_mdoc] eu.europa.ec.eudi.pid.1 (19 claims)\n", "INFO", "[mso_mdoc] eu.europa.ec.eudi.pid.1 (19 claims)", ""},
		{"[VCI] Token response:\n{\n  \"a\": 1\n}\n", "INFO", "Token response:\n{\n  \"a\": 1\n}", "VCI"},
	}
	for _, tc := range cases {
		var buf bytes.Buffer
		fmt.Fprint(NewWriter(&buf), tc.line)
		recs := records(t, buf.String())
		if len(recs) != 1 {
			t.Fatalf("%q: %d records, want 1", tc.line, len(recs))
		}
		rec := recs[0]
		if rec["level"] != tc.level || rec["msg"] != tc.msg {
			t.Errorf("%q: level=%v msg=%q, want %s %q", tc.line, rec["level"], rec["msg"], tc.level, tc.msg)
		}
		if component, _ := rec["component"].(string); component != tc.component {
			t.Errorf("%q: component=%q, want %q", tc.line, component, tc.component)
		}
	}
}

func TestWriterSkipsSeparators(t *testing.T) {
	var buf bytes.Buffer
	w := NewWriter(&buf)
	fmt.Fprintln(w, "───────────────────────────────────────")
	fmt.Fprintln(w)
	if buf.Len() != 0 {
		t.Errorf("separator lines produced records: %s", buf.String())
	}
}

func TestRedirectProcessOutput(t *testing.T) {
	var buf bytes.Buffer
	restore, err := RedirectProcessOutput(&buf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("  Server:      http://localhost:8085")
	fmt.Fprintln(os.Stderr, "warning: stderr line")
	color.New(color.FgYellow).Printf("  Verifier: x\n")
	log.Printf("[VP] Encrypting response")
	restore()

	var msgs []string
	for _, rec := range records(t, buf.String()) {
		msgs = append(msgs, fmt.Sprint(rec["level"], " ", rec["msg"]))
	}
	got := strings.Join(msgs, "|")
	for _, want := range []string{"INFO Server: http://localhost:8085", "WARN stderr line", "INFO Verifier: x", "INFO Encrypting response"} {
		if !strings.Contains(got, want) {
			t.Errorf("records %q miss %q", got, want)
		}
	}
	if strings.Contains(buf.String(), "\x1b[") {
		t.Errorf("records carry color escapes: %s", buf.String())
	}
}

func TestParseFormat(t *testing.T) {
	for in, want := range map[string]string{"": FormatText, "text": FormatText, "JSON": FormatJSON} {
		if got, err := ParseFormat(in); err != nil || got != want {
			t.Errorf("ParseFormat(%q) = %q, %v", in, got, err)
		}
	}
	if _, err := ParseFormat("xml"); err == nil {
		t.Error("ParseFormat(xml) accepted")
	}
}
