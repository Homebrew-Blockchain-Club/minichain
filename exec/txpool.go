package exec

import "github.com/Homebrew-Blockchain-Club/minichain/entity"

type TxPool struct {
}

// 向合约池放入新的交易数据
func (*TxPool) Insert(entity.Transaction) {

}

// 从合约池中取出第一个交易，并从合约池中弹出该交易
func (*TxPool) Poll() entity.Transaction {
	return entity.Transaction{}
}

func (*TxPool) Length() int {
	return 0
}

func (*TxPool) IsFull() bool{
	return false
} 