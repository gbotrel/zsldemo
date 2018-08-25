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

package js

import (
	"encoding/hex"
	"strings"

	ethereum "github.com/gbotrel/gopherjs-eth"
)

type PrivateBank struct {
	contract *ethereum.Contract
	Address  string
}

const (
	EventShielding   = "LogShielding"
	EventUnshielding = "LogUnshielding"
	EventTransfer    = "LogShieldedTransfer"
	EventNullifier   = "LogSpendNullifier"
	EventNewNote     = "LogNewNote"
)

func DeployPrivateBank(eth *ethereum.Ethereum, opts ethereum.CallOpts) (*PrivateBank, error) {
	var err error
	toReturn := &PrivateBank{}
	opts.Data = PrivateBankBin
	toReturn.Address, toReturn.contract, err = eth.DeployContract(PrivateBankABI, opts)
	return toReturn, err
}

func NewPrivateBank(eth *ethereum.Ethereum, address string) *PrivateBank {
	return &PrivateBank{contract: eth.Contract(PrivateBankABI, address), Address: address}
}

func (bank *PrivateBank) SubscribeToEvents(eth *ethereum.Ethereum, chEvents chan *ethereum.Event) {
	bank.contract.SubscribeToEvent(EventShielding, chEvents)
	bank.contract.SubscribeToEvent(EventUnshielding, chEvents)
	bank.contract.SubscribeToEvent(EventTransfer, chEvents)
	bank.contract.SubscribeToEvent(EventNullifier, chEvents)
	bank.contract.SubscribeToEvent(EventNewNote, chEvents)
}

func (bank *PrivateBank) BroadcastNote(opts ethereum.CallOpts, pk, rho []byte, value uint64) (string, error) {
	toReturn, err := bank.contract.Call("broadcastNote", opts,
		hexEncode0x(pk),
		hexEncode0x(rho),
		value)
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) Shield(opts ethereum.CallOpts, proof, sendNullifer, noteCommitment []byte) (string, error) {
	toReturn, err := bank.contract.Call("shield", opts,
		hexEncode0x(proof),
		hexEncode0x(sendNullifer),
		hexEncode0x(noteCommitment))
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) Unshield(opts ethereum.CallOpts, proof, spendNullifer []byte, cm, treeRoot string, value uint64) (string, error) {
	hexProof := hexEncode0x(proof)
	hexNullifier := hexEncode0x(spendNullifer)
	toReturn, err := bank.contract.Call("unshield", opts, hexProof, hexNullifier, cm, treeRoot, value)
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) Transfer(opts ethereum.CallOpts,
	proof, sendNullifier1, sendNullifier2, spendNullifier1, spendNullifier2, cm1, cm2 []byte,
	treeRoot string) (string, error) {

	toReturn, err := bank.contract.Call("shieldedTransfer", opts,
		hexEncode0x(proof),
		treeRoot,
		hexEncode0x(spendNullifier1),
		hexEncode0x(spendNullifier2),
		hexEncode0x(sendNullifier1),
		hexEncode0x(sendNullifier2),
		hexEncode0x(cm1),
		hexEncode0x(cm2))
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) TreeRoot() (string, error) {
	root, err := bank.contract.Call("root", ethereum.CallOpts{})
	if err != nil {
		return "", err
	}
	return root.String(), nil
}

func (bank *PrivateBank) GetWitness(cm string) (uint64, []string, string, error) {
	witness, err := bank.contract.Call("getWitness", ethereum.CallOpts{}, cm)
	if err != nil {
		return 0, nil, "", err
	}
	treeIndex := witness.Index(0).Uint64()
	treePath := strings.Split(witness.Index(1).String(), ",")
	treeRoot := witness.Index(2).String()
	return treeIndex, treePath, treeRoot, nil
}

func (bank *PrivateBank) ShieldedTransferCount() (string, error) {
	toReturn, err := bank.contract.Call("shieldedTransferCount", ethereum.CallOpts{})
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) ShieldingCount() (string, error) {
	toReturn, err := bank.contract.Call("shieldingCount", ethereum.CallOpts{})
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) UnshieldingCount() (string, error) {
	toReturn, err := bank.contract.Call("unshieldingCount", ethereum.CallOpts{})
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func (bank *PrivateBank) TotalSupply() (string, error) {
	toReturn, err := bank.contract.Call("totalSupply", ethereum.CallOpts{})
	if err != nil {
		return "", err
	}
	return toReturn.String(), nil
}

func hexEncode0x(input []byte) string {
	return "0x" + hex.EncodeToString(input)
}
