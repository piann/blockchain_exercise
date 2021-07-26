package blockchain

import (
	"fmt"
	"sync"

	"github.com/piann/coin_101/db"
	"github.com/piann/coin_101/utils"
)

const (
	defaultDifficulty          int = 2
	difficultyRecheckInterval  int = 5
	expectedMiningTimePerBlock int = 1
	allowedRange               int = 1
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"Height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecaculatedBlock := allBlocks[difficultyRecheckInterval-1]
	minuteGap := (newestBlock.Timestamp - lastRecaculatedBlock.Timestamp) / 60
	expectedGap := difficultyRecheckInterval * expectedMiningTimePerBlock
	if minuteGap < expectedGap-allowedRange {
		return b.CurrentDifficulty + 1
	} else if minuteGap > expectedGap+allowedRange {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) getDifficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyRecheckInterval == 0 {
		// caculate again !
		return b.recalculateDifficulty() // temp logic
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{Height: 0}
			// find check point on DB
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				b.restore(checkpoint)

			}

		})
	}
	fmt.Println(b.NewestHash)
	return b
}
