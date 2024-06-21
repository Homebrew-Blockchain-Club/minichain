package comm

import (
	"log"
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

type MockController struct {
	Transactions []entity.Transaction
	Blocks       []ds.Block
}

func (m *MockController) AddTransaction(tx entity.Transaction) {
	m.Transactions = append(m.Transactions, tx)
	log.Println("Mock AddTransaction called with:", tx)
}

func (m *MockController) AddBlock(block ds.Block) {
	m.Blocks = append(m.Blocks, block)
	log.Println("Mock AddBlock called with:", block)
}

func (m *MockController) QueryAccount() entity.Account {
	return entity.Account{}
}

func TestReceive(t *testing.T) {

	// 创建测试交易
	tx := entity.Transaction{}
	txData := typeconv.ToBytes(tx)
	comm := NewCommunicator()

	pkg := Package{
		Type: "transaction",
		Data: txData,
	}

	// 创建MockController
	mockController := &MockController{}

	// 调用receive函数
	comm.Receive(pkg, mockController)

	// 检查MockController是否接收到正确的交易
	if len(mockController.Transactions) != 1 {
		t.Fatalf("Expected 1 transaction, got %d", len(mockController.Transactions))
	}

}
