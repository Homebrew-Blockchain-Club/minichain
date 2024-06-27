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
		newBlock.Header.Prev = hasher.Hash(typeconv.ToBytes(*preBlock))
	}

	// 初始化
	tranTree := ds.NewMPT()
	reciTree := ds.NewMPT()
	var stateTree *ds.MPT
	if preBlock != nil {
		stateTree = ds.NewMPTFromPrevious(preBlock.Header.StateRoot)
	} else {
		stateTree = ds.NewMPT()
	}

	// 处理交易
	for _, tx := range txs {
		result := entity.Result{
			Hash:    hasher.Hash(typeconv.ToBytes(tx)),
			State:   "FAIL",
			Message: "交易验证失败",
		}

		if entity.Verify(tx) {
			// 更新状态树
			res, err := updateStateTree(stateTree, tx)
			if err == nil {
				result = res
			} else {
				result.Message = err.Error()
			}
			newBlock.Transactions = append(newBlock.Transactions, tx)

			// 记录交易树
			tranTree.Update(typeconv.ToHex(typeconv.ToBytes(tx)), typeconv.ToBytes(tx))
		}

		// 记录收据树
		reciTree.Update(typeconv.ToHex(typeconv.ToBytes(tx)), typeconv.ToBytes(result))
	}

	// 提交树
	newBlock.Header.TransactionRoot = tranTree.Commit()
	newBlock.Header.RecipientRoot = reciTree.Commit()
	newBlock.Header.StateRoot = stateTree.Commit()

	// 挖掘区块
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
	for _, tx := range newBlock.Transactions {
		if !entity.Verify(tx) {
			return false
		}
	}
	// 初始化
	tranTreeCmp := ds.NewMPT()
	reciTreeCmp := ds.NewMPT()
	var stateTreeCmp *ds.MPT
	if prevBlock != nil {
		stateTreeCmp = ds.NewMPTFromPrevious(prevBlock.Header.StateRoot)
	} else {
		stateTreeCmp = ds.NewMPT()
	}

	// 处理交易
	for _, tx := range newBlock.Transactions {
		result := entity.Result{
			Hash:    hasher.Hash(typeconv.ToBytes(tx)),
			State:   "FAIL",
			Message: "交易验证失败",
		}

		if entity.Verify(tx) {
			// 更新状态树
			res, err := updateStateTree(stateTreeCmp, tx)
			if err == nil {
				result = res
			} else {
				result.Message = err.Error()
			}
			//newBlock.Transactions = append(newBlock.Transactions, tx)

			// 记录交易树
			tranTreeCmp.Update(typeconv.ToHex(typeconv.ToBytes(tx)), typeconv.ToBytes(tx))
		}

		// 记录收据树
		reciTreeCmp.Update(typeconv.ToHex(typeconv.ToBytes(tx)), typeconv.ToBytes(result))
	}
	tranTreeCmp.Committed = true
	reciTreeCmp.Committed = true
	stateTreeCmp.Committed = true
	// 提交树
	//newBlock.Header.TransactionRoot = tranTree.Commit()
	//newBlock.Header.RecipientRoot = reciTree.Commit()
	//newBlock.Header.StateRoot = stateTree.Commit()

	//验证工作量证明
	blockHash := hasher.Hash(typeconv.ToBytes(newBlock.Header))
	if !isValidHash(blockHash) {
		return false
	}
	//验证交易根哈希
	tranTree := ds.GetMPT(newBlock.Header.TransactionRoot)
	if tranTree == nil || !bytes.Equal(hasher.Hash(typeconv.ToBytes(*tranTreeCmp)), newBlock.Header.TransactionRoot) || !tranTreeCmp.Proof() {
		return false
	}
	//验证收据根哈希
	reciTree := ds.GetMPT(newBlock.Header.RecipientRoot)
	if reciTree == nil || !bytes.Equal(hasher.Hash(typeconv.ToBytes(*reciTreeCmp)), newBlock.Header.RecipientRoot) || !reciTreeCmp.Proof() {
		return false
	}
	//验证状态根哈希
	stateTree := ds.GetMPT(newBlock.Header.StateRoot)
	if stateTree == nil || !bytes.Equal(hasher.Hash(typeconv.ToBytes(*stateTreeCmp)), newBlock.Header.StateRoot) || !stateTreeCmp.Proof() {
		return false
	}
	//验证所有交易的合法性
	newBlock.Header.TransactionRoot = tranTreeCmp.Commit()
	newBlock.Header.RecipientRoot = reciTreeCmp.Commit()
	newBlock.Header.StateRoot = stateTreeCmp.Commit()
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
	// 查询发送者账户
	fromAccountBytes := stateTree.Query(tx.From)
	var fromAccount entity.Account
	typeconv.FromBytesInto(fromAccountBytes, &fromAccount)

	// 查询接收者账户
	toAccountBytes := stateTree.Query(tx.To)
	var toAccount entity.Account
	typeconv.FromBytesInto(toAccountBytes, &toAccount)

	// 检查余额
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

	// 写回状态树
	stateTree.Update(typeconv.ToHex(tx.From), typeconv.ToBytes(fromAccount))
	stateTree.Update(typeconv.ToHex(tx.To), typeconv.ToBytes(toAccount))

	return entity.Result{
		Hash:    hasher.Hash(typeconv.ToBytes(tx)),
		State:   "SUCCESS",
		Message: "",
	}, nil
}
