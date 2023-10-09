package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain is a slice of blocks representing the entire blockchain.
var Blockchain []Block

// CreateHash calculates the hash of a block based on its content.
func CreateHash(b Block) string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// NewBlock creates a new block with a random nonce between 1 and 1000 (inclusive)
// and appends it to the blockchain.
func NewBlock(transaction string, previousHash string) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time
	nonce := rand.Intn(1000) + 1     // Generate a random nonce between 1 and 1000
	newBlock := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	newBlock.Hash = CreateHash(*newBlock)
	Blockchain = append(Blockchain, *newBlock)
}

// DisplayBlocks prints all blocks in the blockchain.
func DisplayBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println()
	}
}

func main() {
	// Create the genesis block (the first block).
	genesisBlock := Block{
		Transaction:  "Genesis Transaction(First Transaction)",
		Nonce:        0,
		PreviousHash: "",
	}
	genesisBlock.Hash = CreateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	// Read user input for transactions and create new blocks.
	for {
		var transaction string
		fmt.Print("Enter a transaction (or 'q' to quit): ")
		fmt.Scanln(&transaction)

		if transaction == "q" {
			break
		}

		NewBlock(transaction, Blockchain[len(Blockchain)-1].Hash)
	}

	// Display all blocks in the blockchain.
	DisplayBlocks()

	// Wait for user to press Enter before exiting.
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
