package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

// This is a Block representing a single block in blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// The entire blockchain.
var Blockchain []Block

// Calculates the hash of a block
func CreateHash(b Block) string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Creating  new block with a random nonce between 1 and 1000 (limit set) and then adding it to blockchain
func NewBlock(transaction string, previousHash string) {
	// Seed generating random number with the current time
	rand.Seed(time.Now().UnixNano())
	//nonce limit is set here
	nonce := rand.Intn(1000) + 1

	newBlock := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}

	newBlock.Hash = CreateHash(*newBlock)
	Blockchain = append(Blockchain, *newBlock)
}

// Display all blocks in the blockchain.
func Display() {
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
	// Generation of first block i.e 1st one.
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
		fmt.Print("Enter a transaction ('q' to quit): ")
		fmt.Scanln(&transaction)

		if transaction == "q" {
			break
		}

		NewBlock(transaction, Blockchain[len(Blockchain)-1].Hash)
	}

	Display()

	//user press enter to exit the code
	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}
