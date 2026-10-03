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

package demorp

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/dominikschlosser/eudi-dev/v2/internal/mock"
	"github.com/dominikschlosser/eudi-dev/v2/internal/wallet"
)

// The EUDI PID and the German PID that extends it both answer a base PID query, in
// both formats.
func twoPIDWallet(t *testing.T) *wallet.Wallet {
	t.Helper()
	holderKey, err := mock.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	issuerKey, err := mock.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	w := wallet.New(holderKey, issuerKey, true)
	if err := w.GenerateProtectedDefaults(); err != nil {
		t.Fatalf("generating the PID baseline: %v", err)
	}
	return w
}

func customAnswers(t *testing.T, status map[string]any) []any {
	t.Helper()
	if status["status"] != "verified" {
		t.Fatalf("status = %v, want verified (error %v, checks %v)", status["status"], status["error"], status["checks"])
	}
	claims, _ := status["claims"].(map[string]any)
	answers, ok := claims["cred_0"].([]any)
	if !ok {
		t.Fatalf("cred_0 = %v, want an array of claim sets", claims["cred_0"])
	}
	return answers
}

func TestVerifierCustomRequestMultipleSDJWT(t *testing.T) {
	_, ts := serveDemoStack(t, twoPIDWallet(t))

	id := createCustom(t, ts, `{"type":"custom","credentials":[{"format":"dc+sd-jwt","vct":"`+PIDVCT+`","claims":[["given_name"]],"multiple":true}]}`)
	answers := customAnswers(t, getJSONFrom(t, ts.URL+"/verifier/api/requests/"+id))
	if len(answers) != 2 {
		t.Fatalf("verified %d presentations, want 2", len(answers))
	}
	for i, a := range answers {
		if claims, _ := a.(map[string]any); claims["given_name"] == nil {
			t.Errorf("presentation %d discloses %v, want given_name", i, a)
		}
	}
}

func TestVerifierCustomRequestMultipleMDoc(t *testing.T) {
	_, ts := serveDemoStack(t, twoPIDWallet(t))

	id := createCustom(t, ts, `{"type":"custom","credentials":[{"format":"mso_mdoc","doctype":"`+PIDDocType+`","claims":[["`+PIDDocType+`","given_name"]],"multiple":true}]}`)
	if answers := customAnswers(t, getJSONFrom(t, ts.URL+"/verifier/api/requests/"+id)); len(answers) != 2 {
		t.Fatalf("verified %d DeviceResponses, want 2", len(answers))
	}
}

// The signed request object carries the flag, so the wallet sees it.
func TestVerifierCustomRequestMultipleReachesTheQuery(t *testing.T) {
	d, ts := serveDemoStack(t, twoPIDWallet(t))

	created := postJSONTo(t, ts.URL+"/verifier/api/requests", `{"type":"custom","credentials":[{"format":"dc+sd-jwt","vct":"`+PIDVCT+`","multiple":true},{"format":"dc+sd-jwt","vct":"`+PIDVCT+`"}]}`)
	id, _ := created["id"].(string)
	d.mu.Lock()
	jar := d.requests[id].requestObject
	d.mu.Unlock()
	parts := strings.Split(jar, ".")
	if len(parts) != 3 {
		t.Fatalf("request object %q is not a compact JWS", jar)
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatal(err)
	}
	var claims struct {
		DCQL struct {
			Credentials []map[string]any `json:"credentials"`
		} `json:"dcql_query"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil {
		t.Fatal(err)
	}
	creds := claims.DCQL.Credentials
	if len(creds) != 2 || creds[0]["multiple"] != true {
		t.Fatalf("dcql_query credentials = %v, want the first with multiple true", creds)
	}
	if _, present := creds[1]["multiple"]; present {
		t.Errorf("the second query carries multiple %v, want it absent", creds[1]["multiple"])
	}
}

// OpenID4VP 1.0 §8.1: without multiple the array holds one presentation.
func TestVerifierCustomRequestRefusesSeveralPresentationsWithoutMultiple(t *testing.T) {
	d := New(twoPIDWallet(t), func() string { return "https://verifier.example" })
	req := &requestState{custom: []customEntry{{queryID: "cred_0", format: "dc+sd-jwt", vct: PIDVCT}}}

	_, checks, err := d.verifyCustomPresentation(req, map[string][]string{"cred_0": {"a~", "b~"}}, &checklist{})
	if err == nil || !strings.Contains(err.Error(), "expected 1 presentation, got 2") {
		t.Fatalf("error = %v, checks %v", err, checks)
	}
}
