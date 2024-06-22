package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/comm"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

func TestReceive(t *testing.T) {

	// 创建测试交易
	tx := entity.Transaction{}
	txData := typeconv.ToBytes(tx)
	Comm := comm.NewCommunicator()

	pkg := comm.Package{
		Type: "transaction",
		Data: txData,
	}

	// 创建MockController
	//mockController := &MockController{}

	// 调用receive函数
	Comm.Receive(pkg)

	// 检查MockController是否接收到正确的交易

}
