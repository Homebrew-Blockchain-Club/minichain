package cli

import "github.com/Homebrew-Blockchain-Club/minichain/entity"

type cli struct {
}

func QueryAccount([]byte) entity.Account {
	return entity.Account{}
}
func CallContract([]byte) []byte {
	return nil
}

// 转账、创建合约、调用合约共用
func NewTransaction(privkey []byte, to []byte, amount []byte, data []byte, gaslimit uint64) {

}
func NewAccount() ([]byte, []byte) {
	return nil, nil
}
