// 本包应当使用/ds内的函数和结构体
package exec

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
)

// 用于核心控制的结构体，上层包应当直接调用这里的函数，其他地方的函数不对包外公开
type Controller struct {
}

func NewController() *Controller {
	return nil
}

// 加入新交易。在交易池未满时放入交易池，交易池已满则进行挖矿，挖矿完成则将新区块返回到Communicator来发布
func (*Controller) AddTransaction(entity.Transaction) {

}

// 加入由其他节点挖出来的新区块。应当先进行验证再加入到本地区块链中，注意加入区块链时是直接对ds进行操作。
// 验证流程：1.满足PoW约束 2.MPT证明通过
func (*Controller) AddBlock(ds.Block) {

}

// 询问账号信息，从/ds中得到账号信息并返回
func (*Controller) QueryAccount() entity.Account {
	return entity.Account{}
}
