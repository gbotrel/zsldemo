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

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

const zslboxURL = "localhost:9000"

var (
	// simulated blockchain
	auth    *bind.TransactOpts
	backend *backends.SimulatedBackend

	// contract tested
	privateBank *PrivateBank
	merkleTree  *ZSLMerkleTree
)

func setupTest(t *testing.T) {
	var err error
	// create a new blockchain simulator
	key, _ := crypto.GenerateKey()
	auth = bind.NewKeyedTransactor(key)
	balance := new(big.Int)
	balance.SetString("2000000000000000000000", 10)
	backend = backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: balance}})
	// deploy a clean instance of the contracts
	_, _, privateBank, err = DeployPrivateBank(auth, backend)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()

	_, _, merkleTree, err = DeployZSLMerkleTree(auth, backend, big.NewInt(29))
	if err != nil {
		t.Fatal(err)
	}

	backend.Commit()
}
