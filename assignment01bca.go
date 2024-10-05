package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

// NewBlock creates a new block and adds it to the blockchain
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Hash:         CalculateHash(transaction, nonce, previousHash),
	}
	Blockchain = append(Blockchain, block)
	fmt.Println("New block added to the blockchain.")
	return &block
}

// ListBlocks prints all blocks in the blockchain
func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock modifies the transaction of a block and recalculates its hash
func ChangeBlock(index int, newTransaction string) {
	if index < 0 || index >= len(Blockchain) {
		fmt.Println("Invalid block index.")
		return
	}

	block := &Blockchain[index]
	block.Transaction = newTransaction
	block.Hash = CalculateHash(block.Transaction, block.Nonce, block.PreviousHash)
	fmt.Println("Block transaction updated.")
}

// VerifyChain verifies the integrity of the blockchain by checking hashes
func VerifyChain() {
	for i := 1; i < len(Blockchain); i++ {
		currentBlock := Blockchain[i]
		previousBlock := Blockchain[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("Blockchain is invalid at block %d\n", i)
			return
		}

		if currentBlock.Hash != CalculateHash(currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash) {
			fmt.Printf("Block %d has been tampered with!\n", i)
			return
		}
	}
	fmt.Println("Blockchain is valid.")
}

// CalculateHash calculates a SHA-256 hash for the given block data
func CalculateHash(transaction string, nonce int, previousHash string) string {
	hashData := transaction + strconv.Itoa(nonce) + previousHash
	hash := sha256.New()
	hash.Write([]byte(hashData))
	return hex.EncodeToString(hash.Sum(nil))
}

// GetLastBlockHash returns the hash of the last block in the blockchain
func GetLastBlockHash() string {
	if len(Blockchain) == 0 {
		return "0" // Genesis block has no previous hash
	}
	return Blockchain[len(Blockchain)-1].Hash
}
