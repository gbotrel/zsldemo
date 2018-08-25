pragma solidity ^0.4.24;

// Original Copyright 2017 Zerocoin Electric Coin Company LLC
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

contract ZSLPrecompile {

    function () public {
        revert();
    }

    function verifyShieldedTransfer(
        bytes memory proof,
        bytes32 treeRoot,
        bytes32 spendNullifier1,
        bytes32 spendNullifier2,
        bytes32 sendNullifier1,
        bytes32 sendNullifier2,
        bytes32 commitment1,
        bytes32 commitment2) public constant returns (bool) {

        bytes32 result;

        // using inline assembly as solidity now forbids call on address with empty code
        assembly {
            // load free memory pointer
            let ptr := mload(0x40)
            let csize := calldatasize()

            // first 4 bytes of calldata = method signature
            calldatacopy(ptr, 4, csize) 

            // call verifyShieldedTransfer precompile
            if iszero(call(gas, 0x8802, 0, ptr, csize, ptr, 0x20)) {
                revert(0, 0)
            }

            // assign output to result
            result := mload(ptr) 

            // set free memory pointer to free space
            mstore(0x40, add(ptr, csize)) 
        }

        return (result[0] == 0x01);
    }

    function verifyShielding(
        bytes memory proof, 
        bytes32 sendNullifier, 
        bytes32 commitment, 
        uint64 value) public constant returns (bool) {

        bytes32 result;

        // using inline assembly as solidity now forbids call on address with empty code
        assembly {
            // load free memory pointer
            let ptr := mload(0x40) 
            let csize := calldatasize()

            // first 4 bytes of calldata = method signature
            calldatacopy(ptr, 4, csize) 

            // call verifyShielding precompile
            if iszero(call(gas, 0x8803, 0, ptr, csize, ptr, 0x20)) {
                revert(0, 0)
            }

            // assign output to result
            result := mload(ptr)

            // set free memory pointer to free space
            mstore(0x40, add(ptr, csize)) 
        }

        return (result[0] == 0x01);
    }

    function verifyUnshielding(
        bytes memory proof,
        bytes32 spendNullifier, 
        bytes32 treeRoot, 
        uint64 value) public constant returns (bool) {

        bytes32 result;

        // using inline assembly as solidity now forbids call on address with empty code
        assembly {
            // load free memory pointer
            let ptr := mload(0x40)
            let csize := calldatasize()

            // first 4 bytes of calldata = method signature
            calldatacopy(ptr, 4, csize)

            // call verifyUnshielding precompile
            if iszero(call(gas, 0x8804, 0, ptr, csize, ptr, 0x20)) {
                revert(0, 0)
            }

            // assign output to result
            result := mload(ptr) 

            // set free memory pointer to free space
            mstore(0x40, add(ptr, csize)) 
        }

        return (result[0] == 0x01);
    }

    function sha256Compress(bytes input) public constant returns (bytes32 result) {
        require(input.length == 64);

        // using inline assembly as solidity now forbids call on address with empty code
        assembly {
            let ptr := mload(0x40)
            // call sha256Compress precompile
            if iszero(call(gas, 0x8801, 0, input, 0x80, ptr, 0x20)) {
                revert(0, 0)
            }

            // assign output to result
            result := mload(ptr)

            // set storage pointer to new space
            mstore(0x40, add(ptr, 0x24)) 
        }
        return result;
    }

}
