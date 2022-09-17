package main

import (
	"fmt"
)

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	//Calculate Hash of a Block

}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer

}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	//change transaction data inside block
}

func ListBlocks(chainHead *Block) {

	//dispaly the data(transaction) inside all blocks

}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"

}

func main() {

	fmt.Println("wtf  is this")

}
