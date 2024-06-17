// 本包应当通过http请求的方式调用/comm给出的web地址
// 接收键盘输入并对应到需要调用的函数
package cli

import "github.com/Homebrew-Blockchain-Club/minichain/entity"

type CLI struct {
}

func NewCLI() *CLI {
	return nil
}

// 根据账户地址返回账户信息 若不存在则返回nil
func (*CLI) QueryAccount([]byte) *entity.Account {
	return nil
}

// 转账、创建合约、调用合约共用的函数
func (*CLI) NewTransaction(privkey []byte, to []byte, amount []byte, data []byte, gaslimit uint64) {

}

// 创建新的公私钥对
func (*CLI) NewAccount() ([]byte, []byte) {
	return nil, nil
}
