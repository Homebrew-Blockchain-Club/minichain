package exec

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

type Controller struct {
}

func (*Controller) AddTransaction(entity.Transaction) {

}
func (*Controller) AddBlock(ds.Block) {

}

func (*Controller) QueryAccount() entity.Account {
	return entity.Account{}
}
