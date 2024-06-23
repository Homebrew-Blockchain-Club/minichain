package exec

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

type AbstractTxPool interface {
	Insert(entity.Transaction)
	Poll() entity.Transaction
	Length() int
	IsFull() bool
}
type AbstractController interface {
	AddTransaction(entity.Transaction)
	AddBlock(ds.Block)
	QueryAccount([]byte) entity.Account
}
