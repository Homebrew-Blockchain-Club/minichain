package exec

import "github.com/Homebrew-Blockchain-Club/minichain/entity"

type TxPool struct {
}

func (*TxPool) Insert(entity.Transaction) {

}
func (*TxPool) Poll() entity.Transaction {
	return entity.Transaction{}
}
