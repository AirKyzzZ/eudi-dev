package wallet

import (
	"bytes"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/fxamacker/cbor/v2"
)

// CIR (EU) 2026/1731 Annex I §§4.1 and 4.2 require protected PID certificate references.
func TestPIDSignatureCertificateReferences(t *testing.T) {
	w := generateTestWallet(t)
	w.IssuerURL = "https://issuer.example"
	if err := w.GenerateDefaultCredentials(nil, ""); err != nil {
		t.Fatal(err)
	}
	srv := NewServer(w, 0, nil)
	for _, cred := range w.Credentials {
		t.Run(cred.Format, func(t *testing.T) {
			var certificateURL string
			var leafDER, digest []byte
			switch cred.Format {
			case "dc+sd-jwt":
				encoded, err := base64.RawURLEncoding.DecodeString(strings.Split(cred.Raw, ".")[0])
				if err != nil {
					t.Fatal(err)
				}
				var header struct {
					X5U string   `json:"x5u"`
					X5T string   `json:"x5t#S256"`
					X5C []string `json:"x5c"`
				}
				if err := json.Unmarshal(encoded, &header); err != nil {
					t.Fatal(err)
				}
				certificateURL = header.X5U
				leafDER, err = base64.StdEncoding.DecodeString(header.X5C[0])
				if err != nil {
					t.Fatal(err)
				}
				digest, err = base64.RawURLEncoding.DecodeString(header.X5T)
				if err != nil {
					t.Fatal(err)
				}
			case "mso_mdoc":
				raw, err := base64.RawURLEncoding.DecodeString(cred.Raw)
				if err != nil {
					t.Fatal(err)
				}
				var document struct {
					IssuerAuth cbor.RawMessage `cbor:"issuerAuth"`
				}
				if err := cbor.Unmarshal(raw, &document); err != nil {
					t.Fatal(err)
				}
				var sign1 []any
				if err := cbor.Unmarshal(document.IssuerAuth, &sign1); err != nil {
					t.Fatal(err)
				}
				var protected map[int64]any
				if err := cbor.Unmarshal(sign1[0].([]byte), &protected); err != nil {
					t.Fatal(err)
				}
				certificateURL, _ = protected[35].(string)
				fingerprint, ok := protected[34].([]any)
				if !ok || len(fingerprint) != 2 || fingerprint[0] != int64(-16) {
					t.Fatal("missing SHA-256 COSE x5t")
				}
				digest = fingerprint[1].([]byte)
				leafDER = sign1[1].(map[any]any)[uint64(33)].([]any)[0].([]byte)
			default:
				t.Fatalf("unexpected default format %s", cred.Format)
			}
			if certificateURL == "" {
				t.Fatal("PID signature has no protected x5u")
			}
			reference, err := url.Parse(certificateURL)
			if err != nil {
				t.Fatal(err)
			}
			response := serverRequest(t, srv, http.MethodGet, reference.Path, "")
			if response.Code != http.StatusOK {
				t.Fatalf("certificate retrieval: %d", response.Code)
			}
			der := response.Body.Bytes()
			if cred.Format == "dc+sd-jwt" {
				block, _ := pem.Decode(der)
				if block == nil {
					t.Fatal("JOSE x5u did not return PEM")
				}
				der = block.Bytes
			}
			if _, err := x509.ParseCertificate(der); err != nil {
				t.Fatal(err)
			}
			want := sha256.Sum256(der)
			if !bytes.Equal(der, leafDER) || !bytes.Equal(digest, want[:]) {
				t.Fatal("certificate reference does not match the signing certificate")
			}
		})
	}
}
