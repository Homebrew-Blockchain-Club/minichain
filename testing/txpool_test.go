package testing

import (
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
		Nonce:       0,
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
	pool.StatDB.Update(typeconv.ToHex(address1.Bytes()), account1Byte)
	pool.StatDB.Update(typeconv.ToHex(address2.Bytes()), account2Byte)

	//str := typeconv.FromBytes[entity.Transaction](pool.StatDB.Query(address1[:]))

	// 创建测试用的交易 41253
	tx1 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 1, Gas: 20}
	tx2 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 2, Gas: 20}
	tx3 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 3, Gas: 5}
	tx4 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 2, Gas: 31}
	tx5 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 3, Gas: 11}

	// 插入交易
	pool.Insert(tx1)
	pool.Insert(tx2)
	pool.Insert(tx3)
	pool.Insert(tx4)
	pool.Insert(tx5)

	// 测试length
	println("length:", pool.Length())
	println("Test:Target(123)")

	println("length(3):", pool.Length())

	// 测试isfull
	tx6 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 4, Gas: 20}
	tx7 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 5, Gas: 20}
	tx8 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 6, Gas: 20}
	tx9 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 7, Gas: 20}
	tx10 := entity.Transaction{From: typeconv.ToHex(address1.Bytes()), Nonce: 8, Gas: 20}

	// 测试queue取交易
	tx11 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 1, Gas: 41}
	tx12 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 5, Gas: 31}
	tx13 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 6, Gas: 21}
	tx14 := entity.Transaction{From: typeconv.ToHex(address2.Bytes()), Nonce: 4, Gas: 21}

	pool.Insert(tx6)
	pool.Insert(tx7)
	pool.Insert(tx8)
	pool.Insert(tx9)
	pool.Insert(tx10)

	println("len：(期待结果8)", pool.Length())
	println("isFull:(期待结果false)", pool.IsFull())

	// 测试将queue中交易插入到pending
	pool.Insert(tx11)

	println("len：(期待结果10)", pool.Length())
	println("isFull:(期待结果true)", pool.IsFull())

	pool.Insert(tx12)
	pool.Insert(tx13)
	pool.Insert(tx14)

	// 为方便演示， 个位数gas为1的为address2
	println("测试取完交易(为方便演示， 个位数gas为1的为address2)：\n当前交易池交易数量为", pool.Length())
	for pool.Length() != 0 {
		tx := pool.Poll()
		println("nonce:", tx.Nonce, "gas:", tx.Gas, "len->", pool.Length())
	}

}
