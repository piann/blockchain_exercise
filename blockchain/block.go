package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/piann/coin_101/db"
	"github.com/piann/coin_101/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"Height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
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

func (b *Block) mine() {

	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n#####\nHash : %s\nTarget : %s\nNonce : %d\n", hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Hash:         "",
		PrevHash:     prevHash,
		Height:       height,
		Difficulty:   Blockchain().getDifficulty(),
		Nonce:        0,
		Transactions: []*Tx{makeCoinbaseTx("Anonymous")},
	}

	block.mine()
	block.persist()
	return block
}
