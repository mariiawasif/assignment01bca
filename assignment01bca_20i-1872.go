// Maria Wasif
// 20i-1872
// Assignment#01
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = block.CreateHash()
	return block
}

func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Adds a new block
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	previousHash := previousBlock.Hash
	newBlock := NewBlock(transaction, nonce, previousHash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) DisplayBlocks() { //Function to print all the blocks
	fmt.Println("----------------- BLOCKCHAIN ---------------------")
	for i, block := range bc.Blocks {
		fmt.Printf("Block #%d\n", i+1)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Block Hash: %s\n", block.Hash)
		fmt.Println("--------------")
	}
}

// Changes the block's transaction
func ChangeBlockTransaction(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Hash = block.CreateHash()
}

// Verifying chain
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		//Condition to verify if the current block's previous hash matches the hash of previous block
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

func main() {

	blockchain := &Blockchain{}
	genesisBlock := NewBlock("Genesis Transaction", 0, "")
	blockchain.Blocks = append(blockchain.Blocks, genesisBlock)

	//
	blockchain.AddBlock("Alina to Hadia", 12345)
	blockchain.AddBlock("Bob to Carol", 67890)

	// Display all blocks in the blockchain.
	blockchain.DisplayBlocks()

	secondBlock := blockchain.Blocks[1]
	ChangeBlockTransaction(secondBlock, "New Transaction: Bob to Dave")

	// Display the updated blockchain.
	fmt.Printf("Updated Blockchain: ")
	blockchain.DisplayBlocks()

	isValid := blockchain.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.") //checks if block is valid or not
	} else {
		fmt.Println("Blockchain is invalid.")
	}

}
