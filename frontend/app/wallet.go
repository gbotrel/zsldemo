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

package app

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

// -------------------------------------------------------------------------------------------------
// Structs
type Wallet struct {
	*js.Object

	keys          *JSArray `js:"keys"`
	notes         *JSArray `js:"notes"`
	busy          bool     `js:"busy"`
	balance       uint64   `js:"balance"`
	selectedNotes []string `js:"selectedNotes"`
}
type FullNote struct {
	note      *Note
	sk        []byte
	nullifier []byte
	status    string // TypeSpent, ...
}

// Seems these can't be in a mixed struct Go / Js obj. See GopherJS Wiki.
var (
	zslClient = NewZSLBoxClient("http://localhost:9001")
	_notes    = make(map[string]*FullNote)
	_keys     = make(map[string][]byte)
)

// -------------------------------------------------------------------------------------------------
// Init
func NewWallet() *Wallet {
	toReturn := &Wallet{Object: newJSObject()}
	toReturn.keys = NewJSArray("pk", "sk")
	toReturn.notes = NewJSArray("commitment", "nullifier", "pk", "value")
	toReturn.busy = false
	toReturn.balance = 0
	toReturn.selectedNotes = make([]string, 0)
	return toReturn
}

func (wallet *Wallet) updateBalance() {
	var b uint64
	for _, v := range _notes {
		if v.status == TypeConfirmed && v.nullifier == nil {
			// add it to balance
			b += v.note.Value
		}
	}
	wallet.balance = b
}

// -------------------------------------------------------------------------------------------------
// GenerateReceiveAddress generates a new key pair (pk, sk)
// where pk can be used as a "receive" address
func (wallet *Wallet) GenerateReceiveAddress() (*ZAddress, error) {
	zAddr, err := zslClient.GetNewAddress(context.Background(), &Void{})
	if err != nil {
		return nil, err
	}
	wallet.addKey(zAddr)
	return zAddr, nil
}

// -------------------------------------------------------------------------------------------------
// Shielding
func (wallet *Wallet) NewShielding(value uint64) (*Shielding, error) {
	// Get new receive address
	zAddr, err := wallet.GenerateReceiveAddress()
	if err != nil {
		return nil, err
	}

	// build our note
	note := &Note{
		Rho:   randomBytes(32),
		Pk:    zAddr.Pk,
		Value: value,
	}

	// create the shielding proof
	shielding, err := zslClient.CreateShielding(context.Background(), note)
	if err != nil {
		return nil, err
	}

	// update our note array
	wallet.addNote(hexEncode0x(shielding.Commitment), &FullNote{note: note, status: TypePending, nullifier: nil, sk: zAddr.Sk})

	return shielding, nil
}

// ConfirmNote marks a note  as mined in the blockchain
func (wallet *Wallet) ConfirmNote(commitment string) error {
	note, ok := _notes[commitment]
	if !ok {
		return errors.New("note with commitment " + commitment + " doesn't exist")
	}
	note.status = TypeConfirmed
	wallet.addNote(commitment, note)

	return nil
}

func (wallet *Wallet) buildShieldedInput(cm string, treeIndex uint64, treePath []string) (*ShieldedInput, error) {
	// get note
	fullNote, ok := _notes[cm]
	if !ok {
		return nil, errors.New("couldn't find note")
	}

	println("build shielded output with sk " + hexEncode0x(fullNote.sk))

	return &ShieldedInput{
		Sk:        fullNote.sk,
		Rho:       fullNote.note.Rho,
		Value:     fullNote.note.Value,
		TreeIndex: treeIndex,
		TreePath:  parseTreePath(treePath),
	}, nil
}

func (wallet *Wallet) updateNote(cm, status string, nullifier []byte) {
	// get note
	fullNote, ok := _notes[cm]
	if !ok {
		panic(errors.New("couldn't find note"))
	}
	fullNote.status = status
	fullNote.nullifier = nullifier
	wallet.addNote(cm, fullNote)
}

// -------------------------------------------------------------------------------------------------
// Unshielding
func (wallet *Wallet) NewUnshielding(cm string, treeIndex uint64, treePath []string) (*Unshielding, uint64, error) {
	shieldedInput, err := wallet.buildShieldedInput(cm, treeIndex, treePath)
	if err != nil {
		return nil, 0, err
	}

	unshielding, err := zslClient.CreateUnshielding(context.Background(), shieldedInput)
	if err != nil {
		return nil, 0, err
	}

	// update our note array
	wallet.updateNote(cm, TypeSpent, unshielding.SpendNullifier)

	return unshielding, shieldedInput.Value, nil
}

// -------------------------------------------------------------------------------------------------
// Transfer
func (wallet *Wallet) NewTransfer(cm1 string, treeIndex1 uint64, treePath1 []string,
	cm2 string, treeIndex2 uint64, treePath2 []string,
	transferAddress string, transferAmount uint64) (*ShieldedTransfer, []*Note, error) {

	var err error

	// decode input to bytes
	outputPk, err := hex.DecodeString(transferAddress[2:])

	// preparing shielded transfer
	request := &ShieldedTransferRequest{
		Inputs:  make([]*ShieldedInput, 2),
		Outputs: make([]*Note, 2),
	}

	if request.Inputs[0], err = wallet.buildShieldedInput(cm1, treeIndex1, treePath1); err != nil {
		return nil, nil, err
	}
	if request.Inputs[1], err = wallet.buildShieldedInput(cm2, treeIndex2, treePath2); err != nil {
		return nil, nil, err
	}

	// difference betwenn balance and what we want to send
	balance := request.Inputs[0].Value + request.Inputs[1].Value
	if balance < transferAmount {
		return nil, nil, errors.New("insufficient funds")
	}

	// output 1 (aka change money) needs a new address
	zAddr, err := wallet.GenerateReceiveAddress()
	if err != nil {
		return nil, nil, err
	}

	// transaction output 1 (aka change money)
	request.Outputs[0] = &Note{
		Rho:   randomBytes(32),
		Pk:    zAddr.Pk,
		Value: balance - transferAmount,
	}

	// transaction output 2 (aka recipient money)
	request.Outputs[1] = &Note{
		Rho:   randomBytes(32),
		Pk:    outputPk,
		Value: transferAmount,
	}

	// call zslbox and create to create our proof
	toReturn, err := zslClient.CreateShieldedTransfer(context.Background(), request)
	if err != nil {
		return nil, nil, err
	}

	// update our note array
	wallet.updateNote(cm1, TypeSpent, toReturn.SpendNullifiers[0])
	wallet.updateNote(cm2, TypeSpent, toReturn.SpendNullifiers[1])

	// add our change money to array
	wallet.addNote(hexEncode0x(toReturn.Commitments[0]), &FullNote{note: request.Outputs[0], status: TypePending, nullifier: nil})

	// if we own the sk corresponding to pk of the second note (i.e we are the payer and the payee)
	if _, ok := _keys[transferAddress]; ok {
		wallet.addNote(hexEncode0x(toReturn.Commitments[1]), &FullNote{note: request.Outputs[1], status: TypePending, nullifier: nil})
	}

	return toReturn, request.Outputs, nil
}

func (wallet *Wallet) ReceivedNote(pk, rho string, value uint64) {
	if sk, ok := _keys[pk]; ok {
		note := &Note{
			Pk:    hexDecode0x(pk),
			Value: value,
			Rho:   hexDecode0x(rho),
		}
		cm, err := zslClient.GetCommitment(context.Background(), note)
		if err != nil {
			println(err)
			return
		}
		commitment := hexEncode0x(cm.Bytes)
		// we know the private key for this note and don't have it yet
		println("adding note with sk " + hexEncode0x(sk))
		toAdd := &FullNote{sk: sk, status: TypeConfirmed, note: note}
		wallet.addNote(commitment, toAdd)
	}
}

// -------------------------------------------------------------------------------------------------
// JSArray and internal struct update functions
func (wallet *Wallet) addNote(commitment string, note *FullNote) {
	defer wallet.updateBalance()
	_notes[commitment] = note
	pk := hexEncode0x(note.note.Pk)
	nullifier := hexEncode0x(note.nullifier)
	wallet.notes.UpdateEntry(commitment, JSCell{Type: note.status, Value: commitment}, JSCell{Value: nullifier}, JSCell{Value: pk}, JSCell{Type: TypeZETH, Value: note.note.Value})
}

func (wallet *Wallet) addKey(zAddr *ZAddress) {
	pk := hexEncode0x(zAddr.Pk)
	sk := hexEncode0x(zAddr.Sk)
	_keys[pk] = zAddr.Sk
	wallet.keys.UpdateEntry(pk, JSCell{Value: pk}, JSCell{Value: sk})
}
