// Copyright 2018 ConsenSys AG
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

package native

import "testing"

func TestAddCommitment(t *testing.T) {
	setupTest(t)
	t.Log("adding a random commitment to the tree")
	cm := randomBytes()
	if _, err := merkleTree.AddCommitment(auth, cm); err != nil {
		t.Fatal(err)
	}
	backend.Commit()

	t.Log("ensuring the tree contains our commitment")
	cmExists, err := merkleTree.CommitmentExists(nil, cm)
	if err != nil {
		t.Fatal(err)
	}

	if !cmExists {
		t.Fatal("commitment should exist after being added")
	}
}
