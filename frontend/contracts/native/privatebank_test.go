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
	"context"
	"encoding/hex"
	"math/big"
	"math/rand"
	"testing"

	"github.com/consensys/zslbox/zsl"
)

// TestVerifyShielding creates a shielding with random input (calling zslbox)
// and verifies it onchain (calling PrivateBank.sol)
func TestVerifyShielding(t *testing.T) {
	setupTest(t)
	// connect to zsl box
	t.Log("connecting to ", zslboxURL)
	client, err := zsl.NewClient(zslboxURL)
	defer client.Close()
	if err != nil {
		t.Fatal(err)
	}

	createShielding(t, client, 2)
}

// TestVerifyUnshielding creates a shielding, unshields it
// and verifies it onchain (calling PrivateBank.sol)
func TestVerifyUnshielding(t *testing.T) {
	setupTest(t)
	// connect to zsl box
	t.Log("connecting to ", zslboxURL)
	client, err := zsl.NewClient(zslboxURL)
	defer client.Close()
	if err != nil {
		t.Fatal(err)
	}

	shielding, addr, note := createShielding(t, client, 2)

	t.Log("getting witness from smart contract")
	treeIndex, treePath, treeRoot, err := privateBank.GetWitness(nil, toArray(shielding.Commitment))
	if err != nil {
		t.Fatal(err)
	}

	shieldedInput := &zsl.ShieldedInput{
		Sk:        addr.Sk,
		Rho:       note.Rho,
		Value:     note.Value,
		TreeIndex: treeIndex.Uint64(),
		TreePath:  parseTreePath(treePath),
	}

	t.Log("computing unshielding proof")
	unshielding, err := client.ZSLBox.CreateUnshielding(context.Background(), shieldedInput)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(("verifying unshielding on chain"))

	auth.Value = big.NewInt(0)
	_, err = privateBank.Unshield(auth, unshielding.Snark, toArray(unshielding.SpendNullifier), toArray(shielding.Commitment), treeRoot, note.Value)
	if err != nil {
		t.Fatal(err)
	}

	backend.Commit()
}

// TestVerifyTransfer creates 2 shielding (input notes) and a shieldedTransfer
// then verifies the proof on chain (calling PrivateBank.sol)
func TestVerifyTransfer(t *testing.T) {
	setupTest(t)

	// connect to zsl box
	t.Log("connecting to ", zslboxURL)
	client, err := zsl.NewClient(zslboxURL)
	defer client.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("creating shieldings as dummy inputs")
	shielding1, addr1, note1 := createShielding(t, client, 2)
	shielding2, addr2, note2 := createShielding(t, client, 5)

	out1addr, err := client.ZSLBox.GetNewAddress(context.Background(), &zsl.Void{})
	if err != nil {
		t.Fatal(err)
	}
	out2addr, err := client.ZSLBox.GetNewAddress(context.Background(), &zsl.Void{})
	if err != nil {
		t.Fatal(err)
	}

	var commitment [32]byte
	copy(commitment[:], shielding1.Commitment)

	treeIndex1, treePath1, treeRoot1, err := privateBank.GetWitness(nil, commitment)
	if err != nil {
		t.Fatal(err)
	}
	copy(commitment[:], shielding2.Commitment)
	treeIndex2, treePath2, _, err := privateBank.GetWitness(nil, commitment)
	if err != nil {
		t.Fatal(err)
	}

	request := &zsl.ShieldedTransferRequest{
		Inputs:  make([]*zsl.ShieldedInput, 2),
		Outputs: make([]*zsl.Note, 2),
	}

	request.Inputs[0] = &zsl.ShieldedInput{
		Sk:        addr1.Sk,
		Rho:       note1.Rho,
		Value:     note1.Value,
		TreeIndex: treeIndex1.Uint64(),
		TreePath:  parseTreePath(treePath1),
	}

	request.Inputs[1] = &zsl.ShieldedInput{
		Sk:        addr2.Sk,
		Rho:       note2.Rho,
		Value:     note2.Value,
		TreeIndex: treeIndex2.Uint64(),
		TreePath:  parseTreePath(treePath2),
	}

	rho1 := randomBytes()
	rho2 := randomBytes()

	request.Outputs[0] = &zsl.Note{
		Rho:   rho1[:],
		Pk:    out1addr.Pk,
		Value: 0,
	}

	request.Outputs[1] = &zsl.Note{
		Rho:   rho2[:],
		Pk:    out2addr.Pk,
		Value: note1.Value + note2.Value,
	}

	t.Log("computing proof")
	shieldedTransfer, err := client.ZSLBox.CreateShieldedTransfer(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}

	// now send this onchain..
	t.Log("sending tx to blockchain")

	auth.Value = big.NewInt(0)
	_, err = privateBank.ShieldedTransfer(auth, shieldedTransfer.Snark, treeRoot1,
		toArray(shieldedTransfer.SpendNullifiers[0]), toArray(shieldedTransfer.SpendNullifiers[1]),
		toArray(shieldedTransfer.SendNullifiers[0]), toArray(shieldedTransfer.SendNullifiers[1]),
		toArray(shieldedTransfer.Commitments[0]), toArray(shieldedTransfer.Commitments[1]))

	if err != nil {
		t.Fatal(err)
	}

	t.Log("commiting tx")
	backend.Commit()

}

func createShielding(t *testing.T, client *zsl.Client, value uint64) (*zsl.Shielding, *zsl.ZAddress, *zsl.Note) {
	// get a new address
	t.Log("getting a new address")
	address, err := client.ZSLBox.GetNewAddress(context.Background(), &zsl.Void{})
	if err != nil {
		t.Fatal(err)
	}

	rho := randomBytes()
	// create a note
	note := &zsl.Note{
		Pk:    address.Pk,
		Rho:   rho[:],
		Value: value,
	}

	// shield it.
	t.Log("shielding created note")
	shielding, err := client.ZSLBox.CreateShielding(context.Background(), note)
	if err != nil {
		t.Fatal(err)
	}

	// now verify it onchain throu the deposit function
	t.Log("calling privateBank.Shield")
	var sendNullifier, commitment [32]byte
	copy(sendNullifier[:], shielding.SendNullifier)
	copy(commitment[:], shielding.Commitment)
	auth.Value = big.NewInt(int64(value) * 1000000000000000000) // 2 tokens = 2 ETH
	// auth.GasLimit = 5000000000

	_, err = privateBank.Shield(auth, shielding.Snark, sendNullifier, commitment)
	if err != nil {
		t.Fatal(err)
	}

	backend.Commit()

	return shielding, address, note
}

func parseTreePath(treePath [][32]byte) [][]byte {
	toReturn := make([][]byte, len(treePath))
	for i, v := range treePath {
		toReturn[i] = make([]byte, 32)
		copy(toReturn[i], v[:])
	}
	return toReturn
}

func toArray(input []byte) [32]byte {
	var toReturn [32]byte
	copy(toReturn[:], input)
	return toReturn
}

func randomBytes() [32]byte {
	var toReturn [32]byte
	if _, err := rand.Read(toReturn[:]); err != nil {
		panic(err)
	}
	return toReturn
}

func TestVerifyShieldingAssembly(t *testing.T) {
	t.Skip("assembly format verification")
	setupTest(t)

	snark, _ := hex.DecodeString("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	nullifier, _ := hex.DecodeString("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	commitment, _ := hex.DecodeString("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")

	auth.Value = big.NewInt(0)

	privateBank.Shield(auth, snark, toArray(nullifier), toArray(commitment))

	t.Log("commiting tx")
	backend.Commit()
}

func TestVerifyTransferAssembly(t *testing.T) {
	t.Skip("assembly format verification")
	setupTest(t)

	snark, _ := hex.DecodeString("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	treeRoot, _ := hex.DecodeString("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	sp1, _ := hex.DecodeString("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
	sp2, _ := hex.DecodeString("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")
	s1, _ := hex.DecodeString("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")
	s2, _ := hex.DecodeString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	c1, _ := hex.DecodeString("1111111111111111111111111111111111111111111111111111111111111111")
	c2, _ := hex.DecodeString("2222222222222222222222222222222222222222222222222222222222222222")

	auth.Value = big.NewInt(0)
	privateBank.ShieldedTransfer(auth, snark, toArray(treeRoot),
		toArray(sp1), toArray(sp2),
		toArray(s1), toArray(s2),
		toArray(c1), toArray(c2))

	t.Log("commiting tx")
	backend.Commit()
}
