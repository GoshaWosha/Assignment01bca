package main

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

func main() {
	for {
		fmt.Println("1: Miner")
		fmt.Println("2: User")
		fmt.Println("3: Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			minerMenu()
		case 2:
			userMenu()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}

func minerMenu() {
	for {
		fmt.Println("1: Add Block")
		fmt.Println("2: Change Block")
		fmt.Println("3: Back to Main Menu")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Add a block
			var transaction string
			var nonce int
			fmt.Print("Enter transaction (e.g., 'bob to alice'): ")
			fmt.Scanln(&transaction)
			fmt.Print("Enter nonce: ")
			fmt.Scanln(&nonce)
			previousHash := GetLastBlockHash()
			NewBlock(transaction, nonce, previousHash)
		case 2:
			// Change a block
			var blockIndex int
			var newTransaction string
			fmt.Print("Enter block index to modify: ")
			fmt.Scanln(&blockIndex)
			fmt.Print("Enter new transaction: ")
			fmt.Scanln(&newTransaction)
			ChangeBlock(blockIndex, newTransaction)
		case 3:
			return
		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}

func userMenu() {
	for {
		fmt.Println("1: List Blocks")
		fmt.Println("2: Verify Blockchain")
		fmt.Println("3: Back to Main Menu")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ListBlocks()
		case 2:
			VerifyChain()
		case 3:
			return
		default:
			fmt.Println("Invalid option, please choose again.")
		}
	}
}
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

func ListBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n\n", block.Hash)
	}
}

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
