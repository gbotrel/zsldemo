//go:generate abigen -sol solidity/PrivateBank.sol --pkg native --out native/privatebank.go
//go:generate echo "generated privatebank.go"
//go:generate cp js/empty.go.js js/privatebank_abibin.go
//go:generate /bin/sh generategojs.sh
//go:generate echo "generated privatebank_abibin.go"
package contracts
