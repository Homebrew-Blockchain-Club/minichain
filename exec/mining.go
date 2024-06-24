package exec

import (
	"bytes"
	"fmt"
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/hasher"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

// 打包多个交易并生成新区块，生成过程中应不断调整区块nonce值以满足PoW约束
func MiningNewBlock(txs []entity.Transaction) ds.Block {
	var newBlock ds.Block
	preBlock := ds.GetTop()
	if preBlock == nil {
		// 创世区块
		newBlock.Header.Prev = nil
	} else {
		newBlock.Header.Prev = typeconv.ToBytes(*preBlock)
	}

	// 初始化交易树、收据树、状态树
	tranTree := ds.NewMPT()
	reciTree := ds.NewMPT()
	stateTree := ds.NewMPTFromPrevious(newBlock.Header.Prev)

	// 处理交易
	for _, tx := range txs {
		result := entity.Result{
			Hash:    hasher.Hash(typeconv.ToBytes(tx)),
			State:   "FAIL",
			Message: "交易验证失败",
		}

		if entity.Verify(tx) {
			// 更新状态树并记录成功状态
			res, err := updateStateTree(stateTree, tx)
			if err == nil {
				result = res
			} else {
				result.Message = err.Error()
			}
			newBlock.Transactions = append(newBlock.Transactions, tx)

			// 记录交易到交易树
			tranTree.Update(typeconv.ToBytes(tx), typeconv.ToBytes(tx))
		}

		// 记录交易结果到收据树
		reciTree.Update(typeconv.ToBytes(tx), typeconv.ToBytes(result))
	}

	// 提交交易树、收据树和状态树
	newBlock.Header.TransactionRoot = tranTree.Commit()
	newBlock.Header.RecipientRoot = reciTree.Commit()
	newBlock.Header.StateRoot = stateTree.Commit()

	// 挖掘区块，调整Nonce以满足工作量证明的要求
	var nonce uint64
	for {
		newBlock.Header.Nonce = nonce
		blockHash := hasher.Hash(typeconv.ToBytes(newBlock.Header))
		if isValidHash(blockHash) {
			break
		}
		nonce++
	}

	return newBlock
}

// 验证来自其他节点的区块是否合法
func ExamineNewBlock(newBlock *ds.Block) bool {
	//验证区块头的前一个区块哈希
	prevBlock := ds.GetTop()
	if prevBlock != nil {
		prevBlockHash := hasher.Hash(typeconv.ToBytes(*prevBlock))
		if !bytes.Equal(prevBlockHash, newBlock.Header.Prev) {
			return false
		}
	}
	//验证工作量证明
	blockHash := hasher.Hash(typeconv.ToBytes(newBlock.Header))
	if !isValidHash(blockHash) {
		return false
	}
	//验证交易根哈希
	tranTree := ds.GetMPT(newBlock.Header.TransactionRoot)
	if tranTree == nil || !tranTree.Proof() {
		return false
	}
	//验证收据根哈希
	reciTree := ds.GetMPT(newBlock.Header.RecipientRoot)
	if reciTree == nil || !reciTree.Proof() {
		return false
	}
	//验证状态根哈希
	stateTree := ds.GetMPT(newBlock.Header.StateRoot)
	if stateTree == nil || !stateTree.Proof() {
		return false
	}
	//验证所有交易的合法性
	for _, tx := range newBlock.Transactions {
		if !entity.Verify(tx) {
			return false
		}
	}

	return true
}
func isValidHash(hash []byte) bool {
	for i := 0; i < ds.DIFFICULTY; i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
func updateStateTree(stateTree *ds.MPT, tx entity.Transaction) (entity.Result, error) {
	// 查询发送者账户的当前状态
	fromAccountBytes := stateTree.Query(tx.From)
	var fromAccount entity.Account
	typeconv.FromBytesInto(fromAccountBytes, &fromAccount)

	// 查询接收者账户的当前状态
	toAccountBytes := stateTree.Query(tx.To)
	var toAccount entity.Account
	typeconv.FromBytesInto(toAccountBytes, &toAccount)

	// 检查余额是否足够
	if fromAccount.Balance < tx.Amount {
		return entity.Result{
			Hash:    hasher.Hash(typeconv.ToBytes(tx)),
			State:   "FAIL",
			Message: "Insufficient balance",
		}, fmt.Errorf("insufficient balance")
	}

	// 更新账户余额
	fromAccount.Balance -= tx.Amount
	toAccount.Balance += tx.Amount

	// 将更新后的账户状态写回状态树
	stateTree.Update(tx.From, typeconv.ToBytes(fromAccount))
	stateTree.Update(tx.To, typeconv.ToBytes(toAccount))

	return entity.Result{
		Hash:    hasher.Hash(typeconv.ToBytes(tx)),
		State:   "SUCCESS",
		Message: "",
	}, nil
}
