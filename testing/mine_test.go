package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/exec"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestMiningNewBlock(t *testing.T) {
	privatekey1, _ := crypto.GenerateKey()
	privatekey2, _ := crypto.GenerateKey()

	//ret["PrivateKey"] = base64.StdEncoding.EncodeToString(crypto.FromECDSA(privatekey))
	//ret["PublicKey"] = base64.StdEncoding.EncodeToString(crypto.CompressPubkey(&privatekey.PublicKey))
	addr1 := crypto.PubkeyToAddress(privatekey1.PublicKey).Bytes()
	addr2 := crypto.PubkeyToAddress(privatekey2.PublicKey).Bytes()
	mpt := ds.NewMPT()
	mpt.Update(typeconv.ToHex(addr1), typeconv.ToBytes(entity.Account{
		Balance: 1000,
	}))
	mptHash := mpt.Commit()
	ds.SetTop(&ds.Block{
		Header: ds.BlockHeader{
			StateRoot: mptHash,
		},
	})
	txs := []entity.Transaction{
		{
			Nonce:    1,
			Gas:      1000,
			GasLimit: 2000,
			From:     addr1,
			To:       addr2,
			Amount:   100,
			Data:     []byte("transaction data"),
		},
		{
			Nonce:    2,
			Gas:      1500,
			GasLimit: 2500,
			From:     addr2,
			To:       addr1,
			Amount:   200,
			Data:     []byte("another transaction"),
		},
	}
	entity.Sign(&txs[0], crypto.FromECDSA(privatekey1))
	entity.Sign(&txs[1], crypto.FromECDSA(privatekey2))
	// 执行挖矿操作
	newBlock := exec.MiningNewBlock(txs)
	if !exec.ExamineNewBlock(&newBlock) {
		t.Errorf("examine failed")
	}
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
			Prev:            nil, // 创世区块
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
