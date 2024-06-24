package testing

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/exec"
	"testing"
)

func TestMiningNewBlock(t *testing.T) {
	// 准备一些交易作为测试数据
	txs := []entity.Transaction{
		{
			Nonce:    1,
			Gas:      1000,
			GasLimit: 2000,
			From:     []byte("sender1"),
			To:       []byte("receiver1"),
			Amount:   100,
			Data:     []byte("transaction data"),
			R:        []byte("R_value"),
			S:        []byte("S_value"),
			V:        []byte("V_value"),
		},
		{
			Nonce:    2,
			Gas:      1500,
			GasLimit: 2500,
			From:     []byte("sender2"),
			To:       []byte("receiver2"),
			Amount:   200,
			Data:     []byte("another transaction"),
			R:        []byte("R_value"),
			S:        []byte("S_value"),
			V:        []byte("V_value"),
		},
		// 添加更多交易作为测试数据
	}

	// 执行挖矿操作
	newBlock := exec.MiningNewBlock(txs)

	// 检查挖矿结果是否符合预期
	if len(newBlock.Transactions) != len(txs) {
		t.Errorf("Expected %d transactions in new block, got %d", len(txs), len(newBlock.Transactions))
	}

	// TODO: 可以添加更多的测试检查点，如验证新区块的哈希、交易树、收据树和状态树等的有效性。
}

func TestExamineNewBlock(t *testing.T) {
	// 创建一个虚拟的区块作为测试数据
	newBlock := &ds.Block{
		Header: ds.BlockHeader{
			Prev:            nil, // 假设这是创世区块
			TransactionRoot: []byte("transaction_root_hash"),
			RecipientRoot:   []byte("recipient_root_hash"),
			StateRoot:       []byte("state_root_hash"),
			Nonce:           12345, // 假设这个区块的Nonce
		},
		Transactions: nil,
	}

	// 执行区块验证操作
	valid := exec.ExamineNewBlock(newBlock)

	// 检查区块验证结果是否符合预期
	if !valid {
		t.Error("New block validation failed")
	}

	// TODO: 可以添加更多的测试检查点，如验证区块头的前一个区块哈希、工作量证明、交易树、收据树和状态树等的有效性。
}
