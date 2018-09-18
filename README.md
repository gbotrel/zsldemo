# ZSLBox DApp demonstrator


## What is it

This repo will get you started running [ZSLBox](https://github.com/consensys/zslbox) alongside an [ethereum node](https://github.com/gbotrel/go-ethereum) and a DApp demonstrating ZCash-like transactions.

![zsldemo](./zsldemo_screenshot.png?raw=true "ZSLDemo")

## What are ZCash transactions

The [ZCash protocol specification](https://github.com/zcash/zips/blob/master/protocol/protocol.pdf) is a must read.

TLDR; ZCash enables private (aka *shielded*) transactions. It uses [Zero Knowledge Proofs](https://z.cash/technology/zksnarks.html) (computed off-chain in a wallet, verified on-chain) to guarantee privacy and mass-conservation of *notes*.

A *note* is a tuple (*pk*, *value*, *rho*) where:

* *pk* is the "paying" key derived from a "spending" key *sk*. *sk* is kept private and is needed to spend a *note*
*  *rho* is randomness used to ensure uniqueness of *nullifiers*
* *value* is the value of the note

ZCash protocol defines 3 operations which produces proofs:

1. **shielding**: *note* --> (proof, commitment) - printing money
2. **unshielding**: (note, sk, merkeProof) --> (proof, nullifier) - burning money
3. **shieldedTransfer**: (shieldedInputs[note, sk, merkleProof], outputs[note]) --> (proof, nullifiers, commitments) - burns 2 notes & creates 2 notes - with Sum(inputValues) == Sum(outputValues). 

For these 3 operations, the principle is the same; one computes the proof offline in a wallet, and submit it on-chain alongside a claim that can be valided by verifying the proof itself.

To complete a **shielding** operation, one would 

1. create (or use) a key pair (pk,sk) 
2. create a note (pk, rho, value) (i.e. print money)
3. compute a proof and a commitment through the **shielding** operation
4. submit the proof, the commitment and the value on-chain for verification. For the **shielding** operation, the proof verification will only guarantee that the submitter owns a note of value v and that it's commitment was computed correctly. It is the responsibility of the blockchain itself (for Ethereum, the smart contract) to authorize or not the submitter to "print money"

The ZCash protocol is elegant and I invite you to read [the specs](https://github.com/zcash/zips/blob/master/protocol/protocol.pdf) to know more, it is time well spent.

The whole privacy point is the **shieldedTransfer** operation. Commitments and nullifiers can't be associated. Transactions are unlinkable and untraceable. The transaction graph is confidential, and linking identities to *notes* is not possible. 

## Previous work

This work and the ZSLBox is based on [Quorum ZSL](https://github.com/jpmorganchase/zsl-q) published in Oct. 2017 by the the ZCash team & JPM. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You'll need to install `go`, `docker` and `docker-compose`. Computing zkSNARKs is memory intensive, so you'll also need to give more RAM (default is 2Gb) to the `docker` daemon. 

### Running

```
go get -u github.com/gbotrel/zsldemo
cd $GOPATH/src/github.com/gbotrel/zsldemo
docker-compose build && docker-compose up
```

You can now access the DApp at `localhost:8001`. `ZSLBox` will cache large files (proving keys) at startup, so the first operations might be slower. 

Current state on MBP 2016:

* **shielding** - success **100%**, average time **3s**
* **unshielding** - success **100%**, average time **17s**
* **shielded transfer** - success **80%**, average time **40s**

### Running the tests

The smart contract tests are written in Go. They run against a simulated blockchain (`github.com/ethereum/go-ethereum/accounts/abi/bind/backends/SimulatedBackend`) which is conveniently using our [ZSL enabled geth](https://github.com/gbotrel/go-ethereum) fork (vendored in `/vendor` through `dep`).

#### Requirement

The tests will hit `ZSLBox`, both for computing the proof and for verifying it through the smart contract. The easiest way to run it is have the same command above (`docker-compose up`) running at the root of the repo. One of the services launched is `ZSLBox`.

```
cd $GOPATH/src/github.com/gbotrel/zsldemo/frontend/contracts/native
go test
```

## Building

### Architecture

#### *WIP TO BE COMPLETED*

3 services, `ZSLBox`, a `geth` running in dev mode and `zsldemo`. 
`zsldemo` serves static files (DApp) that connect to `ZSLBox` through `grpc-web` and to `geth` through `gopherjs-eth` a [GopherJS](https://github.com/gopherjs/gopherjs) wrapper for `web3.js`. 

**Note:** the whole project is coded in Go. 

It is a [GopherJS](https://github.com/gopherjs/gopherjs) experiment, and while few things are frustrating (for example, the 1.5 mb size of the minified JS output), it works surpinsingly well. Go 1.11 is introducing WASM, and I believe in the coming months, we could see a credible dev/prod workflow to develop DApp in Go, whether because it is a native DApp, or because like me, you're allergic to untyped Javascript (`gulp`, `bower`,`npm`,`yarn`,`browserify`,`webpack`,`clojure`,...) and happy with the excellent standard lib and toolchain from go (`go vet`, `go build`, `go test`, `go generate`, `go doc`, `dep`)

If you want to know more or get notified for what's coming next, get in touch (see my profile) or follow ConsenSys & me on Medium. 

