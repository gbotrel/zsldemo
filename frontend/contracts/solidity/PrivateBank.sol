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
import "./ZSLMerkleTree.sol";
import "./ZSLPrecompile.sol";
import "./SafeMath.sol";


/**
 * PrivateBank enables ZCash type of operation ((un)shield ETH and private transfers of "notes")
 */
contract PrivateBank is SafeMath {
    // Depth of the merkle tree decides how many notes this contract can store 2^depth.
    uint constant public ZTOKEN_TREE_DEPTH = 29;

    // Total supply of ZETH
    uint256 public totalSupply;

    // Counters
    uint256 public shieldingCount;
    uint256 public unshieldingCount;
    uint256 public shieldedTransferCount;

    // Merkle tree
    ZSLMerkleTree private tree;

    ZSLPrecompile private zsl;

    // Map of send and spending nullifiers (when creating and consuming shielded notes)
    mapping (bytes32 => uint) private mapNullifiers;

    // Public events to notify listeners
    event LogShielding(address indexed from, uint64 value, bytes32 commitment);
    event LogUnshielding(address indexed to, uint64 value, bytes32 commitment);
    event LogShieldedTransfer(address indexed from, bytes32 cm1, bytes32 cm2);
    event LogSpendNullifier(bytes32 spendNullifier);
    // demo purposes, propagates notes to listeners. 
    event LogNewNote(bytes32 pk, bytes32 rho, uint64 value);

    constructor() public {
        totalSupply = 0;

        // init contracts
        zsl = ZSLPrecompile(new ZSLPrecompile());
        tree = ZSLMerkleTree(new ZSLMerkleTree(ZTOKEN_TREE_DEPTH));
    }

    /**
     * Fallback function to prevent receiving accidental ether
     */
    function () public {
        revert();
    }

    /**
     * demo purposes: broadcast new notes on chain so payee can receive note from payer
     */
    function broadcastNote(bytes32 pk, bytes32 rho, uint64 value) public constant {
        emit LogNewNote(pk, rho, value);
    }

    /**
     * proxy to tree.getWitness
     * @return treeIndex, treePath, treeRoot
     */
    function getWitness(bytes32 commitment) public constant returns (uint, bytes32[], bytes32) {
        return tree.getWitness(commitment);
    }

    /**
     * Shields given amount of ETH (can be 0) and verify that given proof and commitments
     * matches the value
     */
    function shield(bytes proof, bytes32 sendNullifier, bytes32 commitment) public payable {
        uint64 ethValue = uint64(msg.value / 1000000000000000000);
        // require(ethValue >= 1, "incorect eth value"); // or 0 if we allow empty notes 
        require(mapNullifiers[sendNullifier] == 0, "send nullifier exists");     
        require(!tree.commitmentExists(commitment), "commitment exists");
        require(zsl.verifyShielding(proof, sendNullifier, commitment, ethValue), "could not validate shielding");
        tree.addCommitment(commitment);       // will assert if cm has already been added or the tree is full
        mapNullifiers[sendNullifier] = 1;
        emit LogShielding(msg.sender, ethValue, commitment);
        shieldingCount++;
        totalSupply = SafeMath.add(totalSupply, ethValue);
    }

    /**
     * Unshields verify that the given proof is valid and enables calling account to withdraw value ETH
     * /!\ Insecure /!\  @param treeRoot (a witness to unshielding circuit) should be provided by 
     * the smart contract storage, not by the caller. It defeats the purpose of the merkle proof. 
     * TODO: implement a fix-length cache to store treeRoots
     */
    function unshield(bytes proof, bytes32 spendNullifier, bytes32 commitment, bytes32 treeRoot, uint64 value) public {
        require(mapNullifiers[spendNullifier] == 0);
        require(tree.commitmentExists(commitment)); 
        assert(zsl.verifyUnshielding(proof, spendNullifier, treeRoot, value));
        mapNullifiers[spendNullifier] = 1;

        uint256 weiValue = uint256(value) * 1000000000000000000;
        totalSupply = SafeMath.sub(totalSupply, value);
        msg.sender.transfer(weiValue); 

        emit LogSpendNullifier(spendNullifier);
        emit LogUnshielding(msg.sender, value, commitment);
        unshieldingCount++;
    }

    /**
     * shieldedTransfer ensure that 2 notes are spent, and that 2 new notes (cm1, cm2) are valid
     * /!\ Insecure /!\  @param treeRoot (a witness to unshielding circuit) should be provided by 
     * the smart contract storage, not by the caller. It defeats the purpose of the merkle proof. 
     * TODO: implement a fix-length cache to store treeRoots
     */
    function shieldedTransfer(bytes proof, bytes32 treeRoot, 
        bytes32 spendNullifier1, bytes32 spendNullifier2, 
        bytes32 sendNullifier1, bytes32 sendNullifier2, bytes32 cm1, bytes32 cm2 ) public {
        require(mapNullifiers[sendNullifier1] == 0);
        require(mapNullifiers[sendNullifier2] == 0);
        require(mapNullifiers[spendNullifier1] == 0);
        require(mapNullifiers[spendNullifier2] == 0);
        require(!tree.commitmentExists(cm1));
        require(!tree.commitmentExists(cm2));
        assert(zsl.verifyShieldedTransfer(proof, treeRoot,
            spendNullifier1, spendNullifier2,
            sendNullifier1, sendNullifier2,
            cm1, cm2));
        tree.addCommitment(cm1);
        tree.addCommitment(cm2);
        mapNullifiers[sendNullifier1] = 1;
        mapNullifiers[sendNullifier2] = 1;

        // these were missing 
        mapNullifiers[spendNullifier1] = 1;
        mapNullifiers[spendNullifier2] = 1;

        emit LogSpendNullifier(spendNullifier1);
        emit LogSpendNullifier(spendNullifier2);
        emit LogShieldedTransfer(msg.sender, cm1, cm2);
        shieldedTransferCount++;
    }


}
