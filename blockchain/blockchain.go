package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	calcedHash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", calcedHash)
}

func getLastHash() string {
	curBlockchain := GetBlockchain().blocks
	lengthOfBlocks := len(curBlockchain)
	if lengthOfBlocks > 0 {
		return curBlockchain[lengthOfBlocks-1].Hash
	}

	return ""
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}

	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
