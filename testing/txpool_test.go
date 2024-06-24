package testing

import (
	"fmt"
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/exec"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestDefaultTxPool(t *testing.T) {

	// 初始化交易池
	pool := exec.DefualtTxPool{StatDB: ds.NewMPT()}

	// 生成账户
	account1 := entity.Account{
		Nonce:       0,
		Balance:     22,
		StorageRoot: nil,
		Code:        nil,
	}
	account2 := entity.Account{
		Nonce:       1,
		Balance:     223,
		StorageRoot: nil,
		Code:        nil,
	}

	// 生成地址
	privatekey, _ := crypto.GenerateKey()
	publickey := privatekey.PublicKey
	address1 := crypto.PubkeyToAddress(publickey)

	privatekey2, _ := crypto.GenerateKey()
	publickey2 := privatekey2.PublicKey
	address2 := crypto.PubkeyToAddress(publickey2)

	// 存入mpt

	account1Byte := typeconv.ToBytes(account1)
	account2Byte := typeconv.ToBytes(account2)
	fmt.Println(typeconv.ToHex(address1.Bytes()))
	fmt.Println(typeconv.ToHex(address2.Bytes()))
	pool.StatDB.Update(typeconv.ToHex(address1.Bytes()), account1Byte)
	pool.StatDB.Update(typeconv.ToHex(address2.Bytes()), account2Byte)

	//str := typeconv.FromBytes[entity.Transaction](pool.StatDB.Query(address1[:]))

	// 创建测试用的交易 41253
	tx1 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 1, Gas: 20}
	tx2 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 2, Gas: 20}
	tx3 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 3, Gas: 5}
	tx4 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 2, Gas: 30}
	tx5 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 3, Gas: 10}

	// 插入交易
	pool.Insert(tx1)
	pool.Insert(tx2)
	pool.Insert(tx3)
	pool.Insert(tx4)
	pool.Insert(tx5)

	// 检查交易是否正确插入
	f := pool.IsFull()
	println(f)

	println("Test:Target(41253)")
	for i := 0; i < 5; i++ {
		tx := pool.Poll()
		println(tx.Nonce, tx.Gas)
	}

}
