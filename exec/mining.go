package exec

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

// 打包多个交易并生成新区块，生成过程中应不断调整区块nonce值以满足PoW约束
func miningNewBlock([]entity.Transaction) ds.Block {
	return ds.Block{}
}

func examineNewBlock(*ds.Block) bool {
	return false
}
