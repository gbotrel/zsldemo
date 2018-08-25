**Warning: gopherjs-eth is not production ready and should not be deemed secure.**

# gopherjs-eth

`gopherjs-eth` is an experimental [GopherJS](https://github.com/gopherjs/gopherjs) wrapper on web3.js (v0.20.X). 
It is **not** fully compatible with Metamask.

`gopherjs-eth` is not actively maintained.

## Getting Started

### Examples
```
import (
	...
	ethereum "github.com/gbotrel/gopherjs-eth"
)


eth, err := ethereum.Init("http://localhost:8545")

eth.IsConnected()
eth.IsMetamask()
eth.ActiveAccount()
eth.DeployContract(myABI, opts)

```

### Sample project

This [ZSL Demonstrator](https://github.com/...) uses `gopherjs-eth` package.


## License

This project is licensed under the Apache 2 License - see the [LICENSE](LICENSE) file for details


