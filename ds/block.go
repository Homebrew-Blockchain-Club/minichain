package ds

import (
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	storage "github.com/Homebrew-Blockchain-Club/minichain/store"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

const DIFFICULTY = 1

var top *Block

func GetTop() *Block {
	if top != nil {
		return top
	}
	topbyte := storage.Query(typeconv.ToBytes("top"))
	if topbyte == nil {
		top = nil
		return top
	}
	tmp := typeconv.FromBytes[Block](storage.Query(topbyte))
	top = &tmp
	return top
}
func SetTop(b *Block) {
	top = b
	storage.Store(typeconv.ToBytes("top"), typeconv.ToBytes(*b))
}

type BlockHeader struct {
	Nonce           uint64
	Prev            []byte
	StateRoot       []byte
	TransactionRoot []byte
	RecipientRoot   []byte
	Miner           []byte
}
type Block struct {
	Header       BlockHeader
	Transactions []entity.Transaction
}
