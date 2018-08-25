// Package ethereum is a experimental GopherJS wrapper around web3js (v 0.20.x)
package ethereum

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

// Ethereum represents a connection to the blockchain through web3.js
type Ethereum struct {
	jsWeb3 *js.Object
	jsEth  *js.Object
}

type jsResult struct {
	obj *js.Object
	err *js.Object
}

// Init initializes web3.js with a http provider
func Init(nodeURL string) (*Ethereum, error) {
	toReturn := &Ethereum{}

	web3module := js.Global.Get("Web3")
	if web3module == nil {
		return nil, errors.New("Web3 module missing")
	}

	provider := web3module.
		Get("providers").
		Get("HttpProvider").
		New(nodeURL)

	w3 := web3module.New(provider)
	js.Global.Set("web3", w3)

	toReturn.jsWeb3 = w3
	toReturn.jsEth = w3.Get("eth")

	return toReturn, nil
}

// InitMetamask ensure the web3 object is initialized with Metamask provider
// Note: some contract functions are not functional with Metamask
func InitMetamask() (*Ethereum, error) {
	toReturn := &Ethereum{}

	w3metamask := js.Global.Get("web3")
	if w3metamask == nil || !w3metamask.Get("currentProvider").Get("isMetaMask").Bool() {
		return nil, errors.New("web3 not injected, please install Metamask")
	}

	toReturn.jsWeb3 = w3metamask
	toReturn.jsEth = w3metamask.Get("eth")

	return toReturn, nil
}

// IsMetamask returns true if current provider of web3.js object === metamask
func (eth *Ethereum) IsMetamask() bool {
	if isMetaMask := eth.jsWeb3.Get("currentProvider").Get("isMetaMask"); isMetaMask != nil {
		return isMetaMask.Bool()
	}
	return false
}

// ActiveAccount returns web3.eth.accounts[0]
func (eth *Ethereum) ActiveAccount() (string, error) {
	accounts := eth.jsEth.Get("accounts").Interface().([]interface{})
	if accounts == nil || len(accounts) == 0 {
		return "", errors.New("couldn't find active account")
	}
	return accounts[0].(string), nil
}

// DeployContract deploys a new contract and waits for it to be mined
// returns its address , the contract object or an error
func (eth *Ethereum) DeployContract(abi string, opts CallOpts, args ...interface{}) (string, *Contract, error) {
	args = append(args, opts.ToJSObject())

	abiObj := js.Global.Get("JSON").Call("parse", abi)
	txHash := eth.jsEth.Call("contract", abiObj).Call("new", args...).Get("transactionHash").String()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// we need to wait and get tx receipt for txHash
	receipt, err := eth.waitMined(ctx, txHash)
	if err != nil {
		return "", nil, err
	}

	address := receipt.Get("contractAddress").String()

	return address, &Contract{js: eth.jsEth.Call("contract", abiObj).Call("at", address)}, nil
}

// Contract binds an abi and an address to a contract object
func (eth *Ethereum) Contract(abi, address string) *Contract {
	abiObj := js.Global.Get("JSON").Call("parse", abi)
	return &Contract{js: eth.jsEth.Call("contract", abiObj).Call("at", address)}
}

// Balance return the balance (in Eth) of the given account
func (eth *Ethereum) Balance(account string) (*big.Int, error) {
	chResult := make(chan *jsResult)
	eth.jsEth.Call("getBalance", account, func(err, balance *js.Object) {
		chResult <- &jsResult{obj: balance, err: err}
	})
	res := <-chResult
	if res.err != nil {
		return nil, errors.New(res.err.String())
	}

	toReturn := new(big.Int)
	toReturn.SetString(eth.jsWeb3.Call("fromWei", res.obj.String(), "ether").String(), 10)
	return toReturn, nil
}

// ToWei converts ether to wei
func (eth *Ethereum) ToWei(ether *big.Int) *big.Int {
	jsWei := eth.jsWeb3.Call("toWei", ether.String(), "ether")
	toReturn := new(big.Int)
	toReturn.SetString(jsWei.String(), 10)
	return toReturn
}

// FromWei converts wei to ether
func (eth *Ethereum) FromWei(wei *big.Int) *big.Int {
	jsEther := eth.jsWeb3.Call("fromWei", wei.String(), "ether")
	toReturn := new(big.Int)
	toReturn.SetString(jsEther.String(), 10)
	return toReturn
}

// NetworkID returns the networkID web3 object is connected to
func (eth *Ethereum) NetworkID() (int, error) {
	chResult := make(chan *jsResult)
	eth.jsWeb3.Get("version").Call("getNetwork", func(err, network *js.Object) {
		chResult <- &jsResult{obj: network, err: err}
	})
	res := <-chResult
	if res.err != nil {
		return -1, errors.New(res.err.String())
	}
	return res.obj.Int(), nil
}

// IsConnected return  connection status
func (eth *Ethereum) IsConnected() bool {
	return eth.jsWeb3.Call("isConnected").Bool()
}

func (eth *Ethereum) waitMined(ctx context.Context, txHash string) (*js.Object, error) {
	chDone := make(chan *js.Object)
	go func() {
		for {
			// non blocking check if ctx is done
			select {
			case <-ctx.Done():
				chDone <- nil
			default:
			}

			// try to get tx receipt.
			eth.jsEth.Call("getTransactionReceipt", txHash, func(err, res *js.Object) {
				if res != nil {
					chDone <- res // we got it.
					return
				}
			})

			<-time.After(500 * time.Millisecond)
		}
	}()

	if result := <-chDone; result != nil {
		return result, nil
	}
	return nil, errors.New("deadline exceeded")

}
