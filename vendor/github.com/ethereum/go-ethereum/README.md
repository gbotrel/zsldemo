**This is a [ZSL enabled](https://github.com/ConsenSys/zslbox) fork of [geth](https://github.com/ethereum/go-ethereum)**

It exists solely for demonstration purposes and will not be actively maintained.

It adds the following vendored dependencies:

```
govendor fetch github.com/consensys/zslbox
govendor fetch google.golang.org/grpc/...
govendor fetch github.com/golang/protobuf/proto
```

And the file [core/vm/contracts_zsl.go](https://github.com/gbotrel/go-ethereum/blob/zsl/core/vm/contracts_zsl.go), enabling ZSL precompiled contracts (shielding, unshielding and shielded transfers) to be called from a smart contract in solidity.

**Note that the exact same patch can be applied to latest Quorum branch to make it painlessly ZSL enabled.**

## Build

Clone original `geth` repo and add this one as a git remote
```
cd github.com/ethereum/go-ethereum/
git remote add zsl https://github.com/gbotrel/go-ethereum
git fetch --all
git branch --track zsl zsl/zsl
```

Builds and runs as `geth`
```
make geth
```

You can add `ZSLBOX_URL` env variable (defaults to `localhost:9000`) to specify [ZSLBox](https://github.com/ConsenSys/zslbox) URL
```
ZSLBOX_URL=X.X.X.X:PPPP geth ...
```

## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also
included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included
in our repository in the `COPYING` file.