package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func getHash(blockData string) string {
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])

}

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}

func CalculateHash(inputBlock *Block) string {
	//Calculate Hash of a Block
	transactions := inputBlock.transactions
	// joining all the transactions
	var joined_trans = strings.Join(transactions, "")
	var finalHash = joined_trans + inputBlock.prevHash
	return getHash(finalHash)
}

func PrintBlock(chainHead *Block) {

	fmt.Println("Transactions:", chainHead.transactions)
	fmt.Println("BlockHash ;", chainHead.currentHash)

}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {

	//insert new block and return head pointer
	var BlockToInsert = Block{transactions: transactionsToInsert, prevPointer: chainHead, prevHash: chainHead.currentHash}
	BlockToInsert.currentHash = CalculateHash(&BlockToInsert)
	return &BlockToInsert

}

func getTransactionIndex(alltransactions []string, transToFind string) int {

	flag := false
	index := -1
	for i := 0; i < len(alltransactions) && !flag; i++ {

		if alltransactions[i] == transToFind {
			index = i
			flag = true
		}

	}
	return index
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	//change transaction data inside block
	head := chainHead
	var flag = false
	for head != nil && !flag {
		if getTransactionIndex(head.transactions, oldTrans) != -1 {
			changeIndex := getTransactionIndex(head.transactions, oldTrans)
			head.transactions[changeIndex] = newTrans
			head.currentHash = CalculateHash(head)
			flag = true
		}
		head = head.prevPointer
	}
	if !flag {
		fmt.Println("Transaction not found")
	} else {
		fmt.Println("Transaction changed successfully")
	}

}

func ListBlocks(chainHead *Block) {

	//dispaly the data(transaction) inside all blocks
	var currentHead = chainHead
	for currentHead != nil {
		PrintBlock(currentHead)
		currentHead = currentHead.prevPointer
	}

}

func VerifyChain(chainHead *Block) {

	//check whether "Block chain is compromised" or "Block chain is unchanged"
	var chain = chainHead
	var flag = true
	for (chain != nil) && (!flag) {

		if chain.prevPointer != nil {
			if chain.prevHash != chain.prevPointer.currentHash {
				flag = false
			}
		}
		chain = chain.prevPointer
	}
	//check whether "Block chain is compromised" or "Block chain is unchanged"
	if !flag {
		fmt.Println("Block chain is compromised")
	} else {
		fmt.Println("Block chain is unchanged")
	}
}

func main() {

	var genesisBlock = Block{transactions: []string{"a", "b", "c"}, prevPointer: nil, prevHash: "00000"}
	genesisBlock.currentHash = CalculateHash(&genesisBlock)
	fmt.Println(genesisBlock.currentHash)

	// Testing insert block
	var headBlock = &genesisBlock
	newHead := InsertBlock([]string{"f", "y", "p"}, headBlock)
	PrintBlock(newHead)

	// Testing ListBlocks
	ListBlocks(newHead)

	//Testing changeBlock
	ChangeBlock("a", "y", newHead)
	ListBlocks(newHead)

	// Testing verification
	VerifyChain(newHead)

}
