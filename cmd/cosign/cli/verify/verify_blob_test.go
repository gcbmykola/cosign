// Copyright 2022 The Sigstore Authors.
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

package verify

import (
	"testing"
)

func TestSignatures(t *testing.T) {
	sig := "a=="
	b64sig := "YT09"
	tests := []struct {
		description string
		sigRef      string
		shouldErr   bool
	}{
		{
			description: "raw sig",
			sigRef:      sig,
		},
		{
			description: "encoded sig",
			sigRef:      b64sig,
		}, {
			description: "empty ref",
			shouldErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			gotSig, gotb64Sig, err := signatures(test.sigRef)
			if test.shouldErr && err != nil {
				return
			}
			if test.shouldErr {
				t.Fatal("should have received an error")
			}
			if gotSig != sig {
				t.Fatalf("unexpected signature, expected: %s got: %s", sig, gotSig)
			}
			if gotb64Sig != b64sig {
				t.Fatalf("unexpected encoded signature, expected: %s got: %s", b64sig, gotb64Sig)
			}
		})
	}
}
