package blockchain

import (
	"errors"

	"github.com/piann/coin_101/db"
	"github.com/piann/coin_101/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"Height"`
}

var ErrNotFound = errors.New("Block Not Found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	//payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.persist()
	return block
}
