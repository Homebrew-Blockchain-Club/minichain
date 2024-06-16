package ds

import (
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

const DIFFICULTY = 5

var Tail []byte

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
