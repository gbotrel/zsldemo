// Copyright 2018 ConsenSys AG
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/consensys/zslbox/zsl"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
)

// gas cost for zsl operations
const zslGas uint64 = 1

var (
	fail    = []byte{0x00}
	success = []byte{0x01}
)

// gRPC endpoint to ZSLBox
var zslbox zsl.ZSLBoxClient

// init initializes the connection with ZSLBox and registers the zsl precompiled contracts
func init() {
	// ZSLBox URL (default to localhost:9000)
	var zslboxURL string
	if zslboxURL = os.Getenv("ZSLBOX_URL"); zslboxURL == "" {
		zslboxURL = "localhost:9000"
	}

	// connect to zsl box
	zslClient, err := zsl.NewClient(zslboxURL)
	if err != nil {
		fmt.Println("zsl couldn't connect to zslbox", err)
		os.Exit(-1)
	}
	zslbox = zslClient.ZSLBox

	// initialize contracts
	sha256Compress := &sha256Compress{common.BytesToAddress([]byte{0x88, 0x01}), zslGas}
	verifyShieldedTransfer := &verifyShieldedTransfer{common.BytesToAddress([]byte{0x88, 0x02}), zslGas}
	verifyShielding := &verifyShielding{common.BytesToAddress([]byte{0x88, 0x03}), zslGas}
	verifyUnshielding := &verifyUnshielding{common.BytesToAddress([]byte{0x88, 0x04}), zslGas}

	// set contracts
	PrecompiledContractsHomestead[sha256Compress.address] = sha256Compress
	PrecompiledContractsByzantium[sha256Compress.address] = sha256Compress
	PrecompiledContractsHomestead[verifyShieldedTransfer.address] = verifyShieldedTransfer
	PrecompiledContractsByzantium[verifyShieldedTransfer.address] = verifyShieldedTransfer
	PrecompiledContractsHomestead[verifyShielding.address] = verifyShielding
	PrecompiledContractsByzantium[verifyShielding.address] = verifyShielding
	PrecompiledContractsHomestead[verifyUnshielding.address] = verifyUnshielding
	PrecompiledContractsByzantium[verifyUnshielding.address] = verifyUnshielding
}

// -------------------------------------------------------------------------------------------------
// VerifyShielding contract
type verifyShielding struct {
	address common.Address
	gas     uint64
}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyShielding) RequiredGas(input []byte) uint64 {
	return c.gas
}

// Run parse input bytes and calls ZSLBox.VerifyShielding
func (c *verifyShielding) Run(in []byte) ([]byte, error) {
	// ensure input size is what we expect
	if len(in) != 772 {
		err := fmt.Errorf("input must have size of 840 bytes, got %d", len(in))
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	// ignores non-data input
	in = in[32:]
	log.Info("verifyShielding precompiled being called with input", "length", len(in))

	// copy input bytes into our verify request
	verifyRequest := zsl.NewVerifyShieldingRequest()

	copy(verifyRequest.Shielding.SendNullifier, in[:32])
	copy(verifyRequest.Shielding.Commitment, in[32:64])
	// solidty stores uint64 on 32bytes; bytes between 64 and 88 not used.
	verifyRequest.Value = binary.BigEndian.Uint64(in[88:96])
	// 96 +32 (proof size)
	in = in[128:]
	copy(verifyRequest.Shielding.Snark, in[:584])

	// verify the proof
	result, err := zslbox.VerifyShielding(context.Background(), verifyRequest)
	if err != nil {
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	if result.Result {
		return success, nil
	}

	return fail, errors.New("proof is invalid")
}

// -------------------------------------------------------------------------------------------------
// VerifyUnshielding contract
type verifyUnshielding struct {
	address common.Address
	gas     uint64
}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyUnshielding) RequiredGas(input []byte) uint64 {
	return c.gas
}

// Run parse input bytes and calls ZSLBox.VerifyUnshiedling
func (c *verifyUnshielding) Run(in []byte) ([]byte, error) {
	// ensure input size is what we expect
	if len(in) != 772 {
		err := fmt.Errorf("input must have size of 840 bytes, got %d", len(in))
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	// ignores non-data input
	in = in[32:]
	log.Info("verifyUnshielding being called with input", "length", len(in))

	// copy input bytes into our verify request
	verifyRequest := zsl.NewVerifyUnshieldingRequest()

	copy(verifyRequest.SpendNullifier, in[:32])
	copy(verifyRequest.TreeRoot, in[32:64])
	// solidty stores uint64 on 32bytes; bytes between 64 and 88 not used.
	verifyRequest.Value = binary.BigEndian.Uint64(in[88:96])
	in = in[128:]
	copy(verifyRequest.Snark, in[:584])

	// verify the proof
	result, err := zslbox.VerifyUnshielding(context.Background(), verifyRequest)
	if err != nil {
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	if result.Result {
		return success, nil
	}

	return fail, errors.New("proof is invalid")
}

// -------------------------------------------------------------------------------------------------
// VerifyShieldedTransfer contract
type verifyShieldedTransfer struct {
	address common.Address
	gas     uint64
}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *verifyShieldedTransfer) RequiredGas(input []byte) uint64 {
	return c.gas
}

func (c *verifyShieldedTransfer) Run(in []byte) ([]byte, error) {
	// ensure input size is what we expect
	if len(in) != 900 {
		err := fmt.Errorf("input must have size of 900 bytes, got %d", len(in))
		log.Error("zsl", "error", err)
		return []byte{}, err
	}
	// ignores non-data input
	in = in[32:] // ignoring
	log.Info("verifyShieldedTransfer being called with input", "length", len(in))

	// copy input bytes into our verify request
	verifyRequest := zsl.NewVerifyShieldedTransferRequest()

	copy(verifyRequest.TreeRoot, in[:32])
	copy(verifyRequest.ShieldedTransfer.SpendNullifiers[0], in[32:64])
	copy(verifyRequest.ShieldedTransfer.SpendNullifiers[1], in[64:96])
	copy(verifyRequest.ShieldedTransfer.SendNullifiers[0], in[96:128])
	copy(verifyRequest.ShieldedTransfer.SendNullifiers[1], in[128:160])
	copy(verifyRequest.ShieldedTransfer.Commitments[0], in[160:192])
	copy(verifyRequest.ShieldedTransfer.Commitments[1], in[192:224])
	in = in[256:] // 224 + 32 (size of proof)
	copy(verifyRequest.ShieldedTransfer.Snark, in[:584])

	// verify the proof
	result, err := zslbox.VerifyShieldedTransfer(context.Background(), verifyRequest)
	if err != nil {
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	if result.Result {
		return success, nil
	}

	return fail, errors.New("proof is invalid")
}

// -------------------------------------------------------------------------------------------------
// SHA256Compress contract
type sha256Compress struct {
	address common.Address
	gas     uint64
}

func (c *sha256Compress) Run(in []byte) ([]byte, error) {
	if len(in) != 128 {
		err := fmt.Errorf("input must have size of 128 bytes, got %d", len(in))
		log.Error("zsl", "error", err)
		return []byte{}, err
	}

	// skip first bytes
	in = in[32:96]
	result, err := zslbox.Sha256Compress(context.Background(), &zsl.Bytes{Bytes: in})
	if err != nil {
		log.Error("zsl", "error", err)
		return []byte{}, err
	}
	return result.Bytes, nil
}

func (c *sha256Compress) RequiredGas(input []byte) uint64 {
	return c.gas
}
