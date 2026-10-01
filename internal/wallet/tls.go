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
	"net/http"

	"github.com/dominikschlosser/eudi-dev/v2/internal/format"
)

func (w *Wallet) TLSVerification() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.tlsVerificationLocked()
}

func (w *Wallet) tlsVerificationLocked() bool {
	if w.tlsVerify != nil {
		return *w.tlsVerify
	}
	return w.ValidationMode == ValidationModeStrict
}

func (w *Wallet) ConfigureTLS(verify *bool, caPEM []byte) error {
	roots, err := format.TLSRoots(caPEM)
	if err != nil {
		return err
	}
	var override *bool
	if verify != nil {
		value := *verify
		override = &value
	}
	client := format.NewHTTPClient(w.TLSVerification, roots)
	w.mu.Lock()
	previous := w.outboundHTTP
	w.tlsVerify, w.outboundHTTP = override, client
	w.mu.Unlock()
	if previous != nil {
		previous.CloseIdleConnections()
	}
	return nil
}

func (w *Wallet) HTTPClient() *http.Client {
	if w == nil {
		return format.HTTPClientForURL("")
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.outboundHTTP == nil {
		w.outboundHTTP = format.NewHTTPClient(w.TLSVerification, nil)
	}
	return w.outboundHTTP
}

func (w *Wallet) TLSVerificationOverride() *bool {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.tlsVerify == nil {
		return nil
	}
	value := *w.tlsVerify
	return &value
}
