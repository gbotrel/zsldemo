package ethereum

import (
	"encoding/json"
	"errors"
	"math/big"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

// Contract represents a smart contract wrapper
type Contract struct {
	js *js.Object
}

// Event represents an event emitted by the blockchain
type Event struct {
	Address         string
	Event           string
	TransactionHash string
	BlockHash       string
	BlockNumber     int
	LogIndex        int
	Args            map[string]interface{}
}

// CallOpts is an helper to perform transactions
type CallOpts struct {
	Gas  *big.Int
	Wei  *big.Int
	From string
	Data string
}

// UID returns an unique ID for the Event
func (e *Event) UID() string {
	return e.Event + e.Address + strconv.Itoa(e.LogIndex) + strconv.Itoa(e.BlockNumber)
}

// ToJSObject converts CallOpts to a JSObject (json like)
func (opts *CallOpts) ToJSObject() *js.Object {
	jsOpts := js.Global.Get("Object").New()
	if opts.Wei != nil {
		jsOpts.Set("value", opts.Wei.String())
	}
	if opts.Gas != nil {
		jsOpts.Set("gas", opts.Gas.String())
	}
	if len(opts.From) > 0 {
		jsOpts.Set("from", opts.From)
	}
	if len(opts.Data) > 0 {
		jsOpts.Set("data", opts.Data)
	}
	return jsOpts
}

// Call calls method and wait (block) until response
func (c *Contract) Call(method string, opts CallOpts, args ...interface{}) (*js.Object, error) {
	args = append(args, opts.ToJSObject())

	// result channel
	chResult := make(chan *jsResult)

	// async call back
	fn := func(err, res *js.Object) {
		chResult <- &jsResult{obj: res, err: err}
	}
	args = append(args, fn)

	// async call
	c.js.Call(method, args...)

	// wait for response
	res := <-chResult
	if res.err != nil {
		return nil, errors.New(res.err.String())
	}
	return res.obj, nil
}

// SubscribeToEvent will fetch all events from block 0 and watch new ones
// events are sent to channel chEvents
func (c *Contract) SubscribeToEvent(eventName string, chEvents chan *Event) {
	dummy := js.Global.Get("Object").New()
	dummy.Set("fromBlock", 0)
	dummy.Set("toBlock", "latest")
	c.js.Call(eventName, js.Global.Get("Object").New(), dummy, func(err, res *js.Object) {
		if err == nil {
			event := &Event{}
			jsonEvent := js.Global.Get("JSON").Call("stringify", res).String()
			if err1 := json.Unmarshal([]byte(jsonEvent), &event); err1 != nil {
				println("error while parsing json in PrivateBank.sol watcher: " + err1.Error())
			}
			chEvents <- event
		}
	})
}
