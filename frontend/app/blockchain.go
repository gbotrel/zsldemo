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
	ethereum "github.com/gbotrel/gopherjs-eth"
	contracts "github.com/gbotrel/zsldemo/frontend/contracts/js"
	"github.com/gopherjs/gopherjs/js"
)

type Blockchain struct {
	*js.Object

	balance   string `js:"balance"`
	account   string `js:"account"`
	connected bool   `js:"connected"`
	metamask  bool   `js:"metamask"`
	networkID int    `js:"networkID"`

	// private bank stats
	supply         string `js:"supply"`
	cptTransfers   string `js:"cptTransfers"`
	cptShielding   string `js:"cptShielding"`
	cptUnshielding string `js:"cptUnshielding"`

	// arrays
	commitments *JSArray `js:"commitments"`
	nullifiers  *JSArray `js:"nullifiers"`
	events      *JSArray `js:"events"`
}

var privateBank *contracts.PrivateBank

func NewBlockchain() *Blockchain {
	toReturn := &Blockchain{Object: newJSObject()}
	toReturn.balance = "-"
	toReturn.account = "-"
	toReturn.connected = false
	toReturn.metamask = false
	toReturn.networkID = 0
	toReturn.supply = "-"
	toReturn.cptTransfers = "-"
	toReturn.cptShielding = "-"
	toReturn.cptUnshielding = "-"
	toReturn.commitments = NewJSArray("commitment")
	toReturn.nullifiers = NewJSArray("nullifiers")
	toReturn.events = NewJSArray("events")

	return toReturn
}

func (blockchain *Blockchain) AddCommitment(commitment string) {
	blockchain.commitments.UpdateEntry(commitment, JSCell{Value: commitment})
}

func (blockchain *Blockchain) AddNullifier(nullifier string) {
	blockchain.nullifiers.UpdateEntry(nullifier, JSCell{Value: nullifier})
}

func (blockchain *Blockchain) LogEvent(event *ethereum.Event) {

	data := ""
	uid := event.UID()

	switch event.Event {
	case contracts.EventShielding:
		cm := event.Args["commitment"].(string)
		data = "<span class='application__log_type'>shielding</span> <span class='application__log_value'>" + event.Args["value"].(string) + "</span><span class='currency__eth'>ETH</span>"
		data += "<span>&nbsp;-->&nbsp;&nbsp;" + cm + "</span>"
		uid += cm
	case contracts.EventUnshielding:
		cm := event.Args["commitment"].(string)
		data = "<span class='application__log_type'>unshielding&nbsp;</span> "
		data += "<span class='application__log_value'>" + event.Args["value"].(string) + "</span><span class='currency__eth'>ETH</span>"
		data += "<span>&nbsp;<--&nbsp;" + cm + "</span>"
		uid += cm
	case contracts.EventTransfer:
		cm1 := truncateString(event.Args["cm1"].(string), 18)
		cm2 := truncateString(event.Args["cm2"].(string), 18)
		data = "<span class='application__log_transfer'>shielded transfer</span>"
		data += "<span>(?,?) --> (" + cm1 + ",&nbsp;&nbsp;" + cm2 + ")</span>"
		uid += cm1 + cm2
	default:
		return
	}
	addr := "-"
	if e, ok := event.Args["from"]; ok {
		addr = e.(string)
	}
	if e, ok := event.Args["to"]; ok {
		addr = e.(string)
	}
	uid += addr
	blockchain.events.UpdateEntry(uid, JSCell{Value: data})
}
