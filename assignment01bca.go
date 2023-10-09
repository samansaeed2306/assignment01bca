// blockchain.go
package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var blockchain []*block

// NewBlock creates a new block
func NewBlock(transaction string, nonce int, previousHash string) *block {
	newBlock := &block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}

	newBlock.Hash = CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash))
	blockchain = append(blockchain, newBlock)

	return newBlock
}

// DisplayBlocks displays all blocks
func DisplayBlocks() {
	for _, b := range blockchain {
		fmt.Println("Transaction:", b.Transaction)
		fmt.Println("Nonce:", b.Nonce)
		fmt.Println("Previous Hash:", b.PreviousHash)
		fmt.Println("Hash:", b.Hash)
		fmt.Println("--------------")
	}
}

// ChangeBlock modifies a block's transaction
func ChangeBlock(index int, newTransaction string) {
	if index < 0 || index >= len(blockchain) {
		fmt.Println("Invalid block reference")
		return
	}
	blockchain[index].Transaction = newTransaction
	blockchain[index].Hash = CalculateHash(fmt.Sprintf("%s%d%s", newTransaction, blockchain[index].Nonce, blockchain[index].PreviousHash))
}

// VerifyChain checks if the blockchain is valid
func VerifyChain() bool {
	for i := 1; i < len(blockchain); i++ {
		currentBlock := blockchain[i]
		previousBlock := blockchain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}

		currentBlockHash := CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash))
		if currentBlock.Hash != currentBlockHash {
			return false
		}
	}
	return true
}

// CalculateHash calculates the SHA256 hash for a string
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hash[:])
}
