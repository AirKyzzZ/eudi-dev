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

package wallet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/dominikschlosser/eudi-dev/v2/internal/mock"
)

func TestActivityRecordsEncryptedCredentialExchange(t *testing.T) {
	w := generateTestWallet(t)
	issuerKey, err := mock.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	jwk := testEncJWK(t, &issuerKey.PublicKey)
	jwk["kid"] = "issuer-encryption"
	responsePlain := `{"credentials":[{"credential":"issued-credential"}]}`
	responseWire, _, err := EncryptJWE([]byte(responsePlain), &w.HolderKey.PublicKey, "holder", "ECDH-ES", "A128GCM", nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var requestWire, requestPlain string
	issuer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, readErr := io.ReadAll(req.Body)
		if readErr != nil {
			t.Error(readErr)
			rw.WriteHeader(500)
			return
		}
		requestWire = string(body)
		requestPlain, readErr = DecryptCompactJWE(requestWire, issuerKey)
		if readErr != nil {
			t.Error(readErr)
			rw.WriteHeader(400)
			return
		}
		rw.Header().Set("Content-Type", "application/jwt")
		_, _ = io.WriteString(rw, responseWire)
	}))
	defer issuer.Close()
	_, err = w.sendCredentialRequest(credentialRequestAttempt{
		endpoint: issuer.URL, credentialConfigurationID: "pid", metadata: map[string]any{
			"credential_request_encryption": map[string]any{
				"jwks": map[string]any{"keys": []any{jwk}}, "enc_values_supported": []any{"A128GCM"}, "encryption_required": true,
			},
		},
	}, credentialProofs{Type: "jwt", Values: []string{"proof"}})
	if err != nil {
		t.Fatal(err)
	}
	requestEntry := findLogEntry(w.GetLog(), "credential_request")
	responseEntry := findLogEntry(w.GetLog(), "credential_response")
	if requestEntry == nil || requestEntry.Payload == nil || !requestEntry.Payload.Encrypted {
		t.Fatal("request encryption missing")
	}
	if requestEntry.Payload.Wire != requestWire {
		t.Fatal("request wire value differs from the issuer's received body")
	}
	loggedPlain, err := json.Marshal(requestEntry.Payload.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(loggedPlain) != requestPlain {
		t.Fatalf("request plaintext differs: %s != %s", loggedPlain, requestPlain)
	}
	if responseEntry == nil || responseEntry.Payload == nil || !responseEntry.Payload.Encrypted {
		t.Fatal("response encryption missing")
	}
	if responseEntry.Payload.Wire != responseWire || responseEntry.Payload.Body != responsePlain {
		t.Fatal("response differs from the issuer's response")
	}
}

func TestActivityRecordsExactEncryptedPresentation(t *testing.T) {
	w := generateTestWallet(t)
	key, err := mock.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	var received string
	verifier := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, readErr := io.ReadAll(req.Body)
		if readErr != nil {
			t.Error(readErr)
		}
		received = string(body)
		rw.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(rw, `{}`)
	}))
	defer verifier.Close()
	var payload *LogPayload
	_, err = w.SubmitPresentation(&VPTokenMapResult{TokenMap: map[string]string{"pid": "presentation"}}, "", "state", verifier.URL,
		PresentationParams{ResponseMode: "direct_post.jwt", ClientMetadata: map[string]any{"jwks": map[string]any{"keys": []any{testEncJWK(t, &key.PublicKey)}}}},
		func(response *AuthorizationResponseEnvelope) { payload = PresentationLogPayload(response) })
	if err != nil {
		t.Fatal(err)
	}
	if payload == nil || !payload.Encrypted || payload.Wire != received {
		t.Fatal("activity does not contain the transmitted response")
	}
	form, err := url.ParseQuery(received)
	if err != nil {
		t.Fatal(err)
	}
	plain, err := DecryptCompactJWE(form.Get("response"), key)
	if err != nil {
		t.Fatal(err)
	}
	logged, err := json.Marshal(payload.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(logged) != plain {
		t.Fatalf("plaintext mismatch: %s != %s", logged, plain)
	}
}

func TestActivityPayloadSurvivesStorageWithoutChangingDefaultLog(t *testing.T) {
	store := NewWalletStore(t.TempDir())
	w, err := store.LoadOrCreate()
	if err != nil {
		t.Fatal(err)
	}
	w.addProtocolLog("presentation", "presentation_response", "Sending presentation", true,
		map[string]any{"vp_token": "plain"}, &LogPayload{
			Label: "Response", Body: map[string]any{"vp_token": "plain"}, Encrypted: true, Wire: "encrypted-response",
		})
	if err := store.Save(w); err != nil {
		t.Fatal(err)
	}
	w, err = store.LoadOrCreate()
	if err != nil {
		t.Fatal(err)
	}
	s := NewServer(w, 0, nil)
	for _, view := range []string{"", "?view=activity"} {
		rec := httptest.NewRecorder()
		s.mux.ServeHTTP(rec, httptest.NewRequest("GET", "/api/log"+view, nil))
		var entries []LogEntry
		if err := json.Unmarshal(rec.Body.Bytes(), &entries); err != nil {
			t.Fatal(err)
		}
		if len(entries) != 1 {
			t.Fatalf("entries: %s", rec.Body.String())
		}
		if entries[0].Details["vp_token"] != "plain" {
			t.Fatal("existing details changed")
		}
		if view == "" {
			if strings.Contains(rec.Body.String(), `"payload"`) {
				t.Fatal("activity payload leaked into default API")
			}
		} else if entries[0].Payload == nil || !entries[0].Payload.Encrypted || entries[0].Payload.Wire != "encrypted-response" {
			t.Fatalf("lost encrypted payload: %s", rec.Body.String())
		}
	}
}

func TestActivityPreservesRejectedVerifierResponse(t *testing.T) {
	const body = `{"redirect_uri":"/invalid-relative-redirect","diagnostic":"preserve this response"}`
	verifier := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(rw, body)
	}))
	defer verifier.Close()
	key, err := mock.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	for _, mode := range []string{"direct_post", "direct_post.jwt"} {
		t.Run(mode, func(t *testing.T) {
			w := generateTestWallet(t)
			s := NewServer(w, 0, nil)
			result, err := s.deliverAuthorizationError(&AuthorizationRequestParams{
				ClientID: "verifier", ResponseURI: verifier.URL, ResponseMode: mode,
				ClientMetadata: map[string]any{"jwks": map[string]any{"keys": []any{testEncJWK(t, &key.PublicKey)}}},
			}, "access_denied", "declined")
			if err == nil || result != nil {
				t.Fatalf("invalid redirect must still fail: result=%v err=%v", result, err)
			}
			entries := w.GetLog()
			entry := entries[len(entries)-1]
			if entry.Success || entry.Details != nil || entry.Payload == nil || entry.Payload.Body != body {
				t.Fatalf("failed response lost or existing log fields changed: %+v", entry)
			}
		})
	}
}

func TestActivityMetadataCaptureKeepsStreamingDecodeSemantics(t *testing.T) {
	const body = `{"token_endpoint":"https://issuer.example/token"}`
	issuer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Content-Length", "100")
		_, _ = io.WriteString(rw, body)
	}))
	defer issuer.Close()
	payload := &LogPayload{Label: "Response"}
	metadata, err := fetchOAuthMetadata(issuer.URL, payload)
	if err != nil {
		t.Fatalf("a complete JSON value must still decode: %v", err)
	}
	if metadata["token_endpoint"] != "https://issuer.example/token" || payload.Body != body {
		t.Fatalf("metadata or received body lost: metadata=%v payload=%+v", metadata, payload)
	}
}
