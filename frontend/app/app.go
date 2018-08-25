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
	"errors"
	"math/big"
	"strconv"
	"time"

	ethereum "github.com/gbotrel/gopherjs-eth"
	contracts "github.com/gbotrel/zsldemo/frontend/contracts/js"
	"github.com/gopherjs/gopherjs/js"
)

// -------------------------------------------------------------------------------------------------
// App data model
type App struct {
	*js.Object

	// app
	loaded      bool `js:"loaded"`
	hiddenTrick int  `js:"hiddenTrick"`

	// private bank
	bankLoaded  bool   `js:"bankLoaded"`
	bankAddress string `js:"bankAddress"`

	// toast (snackbar)
	errorVisible bool   `js:"errorVisible"`
	toastVisible bool   `js:"toastVisible"`
	toastMessage string `js:"toastMessage"`

	// modal
	modalVisible bool   `js:"modalVisible"`
	modalStatus  string `js:"modalStatus"`
	modalClock   string `js:"modalClock"`

	// shielding
	shieldAmount        uint64 `js:"shieldAmount"`
	shieldingInProgress bool   `js:"shieldingInProgress"`

	// transfer
	transferAmount     uint64 `js:"transferAmount"`
	transferMax        uint64 `js:"transferMax"`
	transferAddress    string `js:"transferAddress"`
	transferInProgress bool   `js:"transferInProgress"`

	// components
	blockchain *Blockchain `js:"blockchain"`
	wallet     *Wallet     `js:"wallet"`

	// blockchain connection
	eth      *ethereum.Ethereum
	chEvents chan *ethereum.Event
}

var ticker *time.Ticker

// -------------------------------------------------------------------------------------------------
// JSArray cell types (mapped with CSS classes)
const (
	TypeETH         = "eth"
	TypeZETH        = "zeth"
	TypePending     = "pending"
	TypeSpent       = "spent"
	TypeConfirmed   = "confirmed"
	TypeShielding   = "shielding"
	TypeUnshielding = "unshielding"
	TypeTransfer    = "transfer"
)

// -------------------------------------------------------------------------------------------------
// Init functions
func NewApp() *App {
	toReturn := &App{
		Object: newJSObject(),
	}
	toReturn.loaded = false
	toReturn.hiddenTrick = 0

	toReturn.bankLoaded = false
	toReturn.bankAddress = ""

	toReturn.errorVisible = false
	toReturn.toastVisible = false
	toReturn.toastMessage = ""

	toReturn.modalVisible = false
	toReturn.modalStatus = ""
	toReturn.modalClock = ""

	toReturn.shieldingInProgress = false
	toReturn.shieldAmount = 0

	toReturn.transferInProgress = false
	toReturn.transferAmount = 0
	toReturn.transferMax = 0
	toReturn.transferAddress = ""

	toReturn.chEvents = make(chan *ethereum.Event)

	toReturn.blockchain = NewBlockchain()
	toReturn.wallet = NewWallet()

	return toReturn
}

func (app *App) OnLoad(event *js.Object) {
	// monitor for changes in web3js (account, network, ..)
	go app.MonitorBlockchain()

	// monitor for blockchain events on contract PrivateBank.sol
	go app.MonitorEvents()
}

func (app *App) updateBlockchain() error {

	if app.eth == nil {
		var err error
		// app.eth, err = ethereum.InitMetamask()
		app.eth, err = ethereum.Init("http://localhost:8545")
		if err != nil {
			return err
		}
	}
	app.blockchain.connected = app.eth.IsConnected()
	if !app.blockchain.connected {
		return errors.New("unable to reach node")
	}
	app.blockchain.metamask = app.eth.IsMetamask()

	activeAccount, err := app.eth.ActiveAccount()
	if err != nil {
		return err
	}
	app.blockchain.account = activeAccount

	networkID, err := app.eth.NetworkID()
	if err != nil {
		return err
	}
	app.blockchain.networkID = networkID
	if app.blockchain.networkID != 1337 {
		return errors.New("network id should be 1337 ")
	}

	balance, err := app.eth.Balance(activeAccount)
	if err != nil {
		return err
	}
	app.blockchain.balance = balance.String()
	return nil
}

// -------------------------------------------------------------------------------------------------
// Monitor web3js for changes
func (app *App) MonitorBlockchain() {
	var err error
	// polling is the recommended way to monitor Metamask.
	for {
		if err = app.updateBlockchain(); err != nil {
			app.loaded = false
			println("error while polling blockchain" + err.Error())
		} else {
			app.loaded = true
		}
		<-time.After(1 * time.Second)
	}
}

// -------------------------------------------------------------------------------------------------
// Monitor private bank events
func (app *App) MonitorEvents() {
	for {
		select {
		case event := <-app.chEvents:
			go func() {
				println("received event: " + event.Event)
				switch event.Event {
				// Shielding
				case contracts.EventShielding:
					app.blockchain.cptShielding, _ = privateBank.ShieldingCount()
					commitment := event.Args["commitment"].(string)
					app.wallet.ConfirmNote(commitment)
					app.blockchain.AddCommitment(commitment)

				// Unshielding
				case contracts.EventUnshielding:
					app.blockchain.cptUnshielding, _ = privateBank.UnshieldingCount()

				// Transfer
				case contracts.EventTransfer:
					app.blockchain.cptTransfers, _ = privateBank.ShieldedTransferCount()
					cm1 := event.Args["cm1"].(string)
					cm2 := event.Args["cm2"].(string)
					app.wallet.ConfirmNote(cm1)
					app.wallet.ConfirmNote(cm2)
					app.blockchain.AddCommitment(cm1)
					app.blockchain.AddCommitment(cm2)

				// New nullifier
				case contracts.EventNullifier:
					nullifier := event.Args["spendNullifier"].(string)
					app.blockchain.AddNullifier(nullifier)
				case contracts.EventNewNote:
					pk := event.Args["pk"].(string)
					value, _ := strconv.Atoi(event.Args["value"].(string))
					rho := event.Args["rho"].(string)

					app.wallet.ReceivedNote(pk, rho, uint64(value))
				}

				// update total supply
				app.blockchain.supply, _ = privateBank.TotalSupply()

				// add it to our logs
				app.blockchain.LogEvent(event)
				app.hiddenTrick++
			}()
		}
	}
}

// -------------------------------------------------------------------------------------------------
// Init private bank
func (app *App) BtnDeployBankClicked() {
	go func() {
		err := app.deployPrivateBank()
		app.bankLoaded = err == nil
	}()

}
func (app *App) BtnBindBankClicked() {
	go func() {
		err := app.bindPrivateBank()
		app.bankLoaded = err == nil
	}()
}

func (app *App) bindPrivateBank() error {
	privateBank = contracts.NewPrivateBank(app.eth, app.bankAddress)
	privateBank.SubscribeToEvents(app.eth, app.chEvents)

	return app.updatePrivateBankStats()
}

func (app *App) deployPrivateBank() error {
	var err error
	opts := ethereum.CallOpts{
		From: app.blockchain.account,
		Gas:  big.NewInt(5000000),
	}

	privateBank, err = contracts.DeployPrivateBank(app.eth, opts)
	if err != nil {
		return err
	}
	privateBank.SubscribeToEvents(app.eth, app.chEvents)

	return app.updatePrivateBankStats()
}

func (app *App) updatePrivateBankStats() error {
	var err error

	if app.blockchain.cptTransfers, err = privateBank.ShieldedTransferCount(); err != nil {
		return err
	}
	if app.blockchain.cptShielding, err = privateBank.ShieldingCount(); err != nil {
		return err
	}
	if app.blockchain.cptUnshielding, err = privateBank.UnshieldingCount(); err != nil {
		return err
	}
	if app.blockchain.supply, err = privateBank.TotalSupply(); err != nil {
		return err
	}
	app.bankAddress = privateBank.Address
	return nil
}

// -------------------------------------------------------------------------------------------------
// Generate new key
func (app *App) BtnGenerateClicked() {
	go func() {
		_, err := app.wallet.GenerateReceiveAddress()
		app.CheckError(err)
		app.hiddenTrick++ // forces view to update
	}()
}

// -------------------------------------------------------------------------------------------------
// Shielded Transfer
func (app *App) BtnTransferClicked() {
	go func() {
		// set max amount and amount to balance of 2 selected notes
		// get input notes
		fullNote1, ok := _notes[app.wallet.selectedNotes[0]]
		if !ok {
			app.CheckError(errors.New("couldn't find selected note"))
		}
		fullNote2, ok := _notes[app.wallet.selectedNotes[1]]
		if !ok {
			app.CheckError(errors.New("couldn't find selected note"))
		}

		app.transferAmount = fullNote1.note.Value + fullNote2.note.Value
		app.transferMax = app.transferAmount

		app.transferInProgress = true
		app.modalVisible = true
	}()
}

func (app *App) FormTransferSubmitted() {
	go func() {
		app.wallet.busy = true

		defer app.CloseModal()
		go app.StartClock()

		cm1 := app.wallet.selectedNotes[0]
		cm2 := app.wallet.selectedNotes[1]

		// get witnesses
		app.modalStatus = "computing zkSnarks to spend notes... (transfer)"
		treeIndex1, treePath1, treeRoot1, err := privateBank.GetWitness(cm1)
		app.CheckError(err)

		treeIndex2, treePath2, treeRoot2, err := privateBank.GetWitness(cm2)
		app.CheckError(err)

		if treeRoot1 != treeRoot2 {
			app.CheckError(errors.New("tree roots are different.."))
		}

		// compute proof
		shieldedTransfer, newNotes, err := app.wallet.NewTransfer(cm1, treeIndex1, treePath1,
			cm2, treeIndex2, treePath2,
			app.transferAddress, app.transferAmount)

		app.CheckError(err)

		app.modalStatus = "sending shielded transfer tx to blockchain..."

		// send tx to blockchain
		opts := ethereum.CallOpts{
			From: app.blockchain.account,
			Gas:  big.NewInt(5000000),
		}
		txHash, err := privateBank.Transfer(opts, shieldedTransfer.Snark,
			shieldedTransfer.SendNullifiers[0], shieldedTransfer.SendNullifiers[1],
			shieldedTransfer.SpendNullifiers[0], shieldedTransfer.SpendNullifiers[1],
			shieldedTransfer.Commitments[0], shieldedTransfer.Commitments[1], treeRoot1)
		app.CheckError(err)
		println("shielded transfer receipt: " + txHash)

		// propagate new notes (demo)
		txHash, err = privateBank.BroadcastNote(opts, newNotes[0].Pk, newNotes[0].Rho, newNotes[0].Value)
		app.CheckError(err)
		txHash, err = privateBank.BroadcastNote(opts, newNotes[1].Pk, newNotes[1].Rho, newNotes[1].Value)
		app.CheckError(err)
	}()
}

// -------------------------------------------------------------------------------------------------
// Unshielding
func (app *App) BtnUnshieldClicked() {
	go func() {
		app.modalVisible = true
		app.wallet.busy = true

		defer app.CloseModal()
		go app.StartClock()

		cm := app.wallet.selectedNotes[0]

		app.modalStatus = "computing zkSnark to spend note... (unshielding)"
		treeIndex, treePath, treeRoot, err := privateBank.GetWitness(cm)
		app.CheckError(err)

		unshielding, value, err := app.wallet.NewUnshielding(cm, treeIndex, treePath)
		app.CheckError(err)

		app.modalStatus = "sending unshielding tx to blockchain..."
		opts := ethereum.CallOpts{
			From: app.blockchain.account,
			Gas:  big.NewInt(5000000),
		}
		txHash, err := privateBank.Unshield(opts, unshielding.Snark, unshielding.SpendNullifier, cm, treeRoot, value)
		app.CheckError(err)
		println("unshielding receipt: " + txHash)
	}()

}

// -------------------------------------------------------------------------------------------------
// Shielding
func (app *App) BtnShieldClicked() {
	app.modalVisible = true
	app.shieldingInProgress = true
}

func (app *App) FormShieldSubmitted() {
	go func() {
		// create our proof
		app.modalStatus = "computing zkSnark for new note... (shielding)"
		app.wallet.busy = true
		defer app.CloseModal()
		go app.StartClock()
		shielding, err := app.wallet.NewShielding(app.shieldAmount)
		app.CheckError(err)
		app.modalStatus = "sending shielding tx to blockchain..."
		opts := ethereum.CallOpts{
			From: app.blockchain.account, // could be nil here.
			Gas:  big.NewInt(5000000),
			Wei:  app.eth.ToWei(big.NewInt(int64(app.shieldAmount))),
		}
		txHash, err := privateBank.Shield(opts, shielding.Snark, shielding.SendNullifier, shielding.Commitment)
		app.CheckError(err)
		println("shielding receipt: " + txHash)
	}()

}

// -------------------------------------------------------------------------------------------------
// Utils
func (app *App) CopyToClipboard(content string) {
	if !copyToClipboard(content) {
		app.CheckError(errors.New("couldn't copy " + content + " to clipboard"))
	} else {
		app.ShowSnackbar(truncateString(content, 10) + " copied to clipboard")
	}
}

func (app *App) CheckError(err error) {
	if err == nil {
		return
	}
	app.errorVisible = true
	defer func() {
		<-time.After(4 * time.Second)
		app.errorVisible = false
	}()
	panic(err.Error())
}

func (app *App) CloseModal() {
	app.wallet.busy = false
	app.modalVisible = false
	app.shieldAmount = 0
	app.transferAmount = 0
	app.transferMax = 0
	app.modalStatus = "-"
	app.shieldingInProgress = false
	app.transferInProgress = false
	app.wallet.selectedNotes = make([]string, 0)
	app.transferAddress = ""
	app.StopClock()
}

func (app *App) ShowSnackbar(message string) {
	app.toastMessage = message
	app.toastVisible = true
	go func() {
		<-time.After(3 * time.Second)
		app.toastVisible = false
		app.toastMessage = ""
	}()
}

func (app *App) StartClock() {
	if ticker != nil {
		ticker.Stop()
	}
	ticker = time.NewTicker(time.Millisecond * 500)
	start := time.Now()

	for _ = range ticker.C {
		elapsed := time.Since(start)
		app.modalClock = elapsed.Truncate(time.Second).String()
	}

}

func (app *App) StopClock() {
	ticker.Stop()
	app.modalClock = ""
}
