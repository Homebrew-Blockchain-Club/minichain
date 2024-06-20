package ds

import (
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

const DIFFICULTY = 5

var top []byte

func GetTop() *Block {
	return nil
}
func SetTop(*Block) {

}

type BlockHeader struct {
	Nonce           uint64
	Prev            []byte
	StateRoot       []byte
	TransactionRoot []byte
	RecipientRoot   []byte
}
type Block struct {
	Header       BlockHeader
	Transactions []entity.Transaction
}
