package main

/**
 *-----------------------------------------
 * Exercice: 1
 *-----------------------------------------
 * Name: 	"Ledger Baker"
 *-----------------------------------------
 * Objectives:
 *
 * - Create a simple structure to register
 * three transactions.
 * - for help on map feature, see: https://blog.golang.org/go-maps-in-action
 */
import (
	"fmt"
	// --------------------------------------
	//TODO 1) : add package
	// "time"
)

// Block defines a node in blockchain structure.
type Block struct {

	// --------------------------------------
	//TODO 2): finish the block structure
	// - add id, lastID, timestamp (string type)
	data Data
}

// Data structure which contains the transaction data.
type Data struct {
	reference string
	quantity  int
	price     float32
}

// Chain structure.
type Chain struct {
	lastKey int
	blocks  map[int]Block
}

// Enrich chain with a add block feature
func (chain *Chain) addBlock(b Block) {

	// --------------------------------------
	//TODO 3): handle :
	// - last id / current id on new block
	// - add block to Chain structure
}

// Create the first block for chain structure.
func genesis() Chain {
	// --------------------------------------
	//TODO 4): review method to init the chain
	return Chain{}
}

// Create a new default block with data.
func generateBlock(reference string, quantity int, price float32) Block {
	return Block{

		// --------------------------------------
		// TODO 5): initialized: id, creation date

		data: Data{reference: reference, quantity: quantity, price: price},
	}
}

// Execute the code in console: "go run bc_exo1.go"
func main() {
	fmt.Println("Simple Block Chain Creation")
	var blockChain = genesis()
	// Add three transactions:
	blockChain.addBlock(generateBlock("croissants", 5, 1.2))
	blockChain.addBlock(generateBlock("pains", 2, 2.3))
	blockChain.addBlock(generateBlock("croissants", 4, 1.77))

	// Display blockchain
	for _, v := range blockChain.blocks {
		fmt.Println("Block: ", v)
	}
}
