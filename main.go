package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timeStamp    time.Time
	pow          int
}

type BlockChain struct {
	genesisBlock Block //  represents the first block added to the blockchain
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	data, err := json.Marshal(b.data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return ""
	}
	blockData := b.previousHash + string(data) + b.timeStamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) BlockChain {
	genesisBlock := Block{hash: "0", timeStamp: time.Now()}
	return BlockChain{
		genesisBlock: genesisBlock,
		chain:        []Block{genesisBlock},
		difficulty:   difficulty,
	}
}

func (b *BlockChain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timeStamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b BlockChain) isvalid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]

		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	blockChain := CreateBlockchain(2)

	blockChain.addBlock("Anna", "John", 100)
	blockChain.addBlock("Adam", "Sarah", 75)

	fmt.Println(blockChain.isvalid())
}
