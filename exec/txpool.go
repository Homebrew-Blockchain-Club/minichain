package exec

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

// MAX_POOL_SIZE PoolSize fixed
const MAX_POOL_SIZE = 10

type TxPool interface {
	Insert(tx entity.Transaction)
	Poll() entity.Transaction
	IsFull() bool
}

type SortedTxsBlk struct {
	Txs      []entity.Transaction
	GasPrice uint64 // equal to Gas in the transaction struct
}

type RowBlks []SortedTxsBlk

func (td RowBlks) Len() int {
	return len(td)
}
func (td RowBlks) Less(i, j int) bool {
	return td[i].GasPrice > td[j].GasPrice

}
func (td RowBlks) Swap(i, j int) {
	td[i], td[j] = td[j], td[i]
}

type Txs []entity.Transaction

func (td Txs) Len() int {
	return len(td)
}
func (td Txs) Less(i, j int) bool {
	return td[i].Nonce < td[j].Nonce

}
func (td Txs) Swap(i, j int) {
	td[i], td[j] = td[j], td[i]
}

type DefualtTxPool struct {
	StatDB   *ds.MPT
	Pendings map[string]RowBlks // map to store pending transactions grouped by nonce
	Queues   map[string]Txs     // queue for out-of-order nonce transactions
	MiniPool Txs                // in order to poll one by one
}

// Insert 向合约池放入新的交易数据
func (pool *DefualtTxPool) Insert(tx entity.Transaction) {
	// Find address to create a new row
	address := common.BytesToAddress(tx.From).Hex()

	rowBlks := pool.Pendings[address]

	accountBytes := pool.StatDB.Query(tx.From)
	account := typeconv.FromBytes[entity.Account](accountBytes)
	if account.Nonce >= tx.Nonce {
		return
	}

	// If the pending row doesn't exist, create a new one
	if len(rowBlks) == 0 {
		newRowBlk := make(RowBlks, 0)
		pool.Pendings[address] = newRowBlk
	}

	// Trying to insert the transaction into the pending list

	if tx.Nonce == account.Nonce+1 || VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {

		if len(rowBlks) == 0 {
			pool.CreateBlk(tx, address)
		} else {
			isInsert := false
			for i := len(rowBlks) - 1; i >= 0; i-- {
				if rowBlks[i].GasPrice == tx.Gas && VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {
					pool.Pendings[address][i].Txs = append(pool.Pendings[address][i].Txs, tx)
					pool.PollFromQueue(tx, address)
					isInsert = true
				}
			}
			if !isInsert {
				pool.CreateBlk(tx, address)
			}
		}

	} else if tx.Nonce > pool.GetPendingLastNonce(address)+1 {
		// If the nonce is out of order, add it to the queue
		queueTxs := pool.Queues[address]
		if len(queueTxs) == 0 {
			pool.Queues = make(map[string]Txs)
			queueTxs = pool.Queues[address]
		}
		queueTxs = append(queueTxs, tx)
		sort.Sort(queueTxs)
	} else if tx.Nonce <= pool.GetPendingLastNonce(address) {
		// If the nonce is less than the last nonce, ignore it
		return
	}
	pool.AllSortedByGas()
}

// CreateBlk Create Blk and append the first new transaction
func (pool *DefualtTxPool) CreateBlk(tx entity.Transaction, address string) {
	newBlk := SortedTxsBlk{
		Txs:      make([]entity.Transaction, 0),
		GasPrice: tx.Gas,
	}
	newBlk.Txs = append(newBlk.Txs, tx)
	pool.Pendings[address] = append(pool.Pendings[address], newBlk)

}

// Poll 从合约池中取出第一个交易，并从合约池中弹出该交易
func (pool *DefualtTxPool) Poll() entity.Transaction {
	tx := pool.MiniPool[0]
	pool.MiniPool = append(pool.MiniPool[1:])

	// Delete the transaction from the pending list
	pool.DeleteFromPending(tx)
	return tx
}

func (pool *DefualtTxPool) Length() int {
	return len(pool.MiniPool)
}

func (pool *DefualtTxPool) AllSortedByGas() {
	allBlk := make(RowBlks, 0)
	for _, rowBlk := range pool.Pendings {
		allBlk = append(allBlk, rowBlk...)
	}
	// Sorted by gas
	sort.Sort(allBlk)

	allTx := make(Txs, 0)
	for _, tx := range allBlk {
		allTx = append(allTx, tx.Txs...)
	}

	// Insert into miniPool

	for i := 0; i < min(MAX_POOL_SIZE, len(allTx)); i++ {
		pool.MiniPool = append(pool.MiniPool, allTx[i])
	}

}

func (pool *DefualtTxPool) IsFull() bool {
	return len(pool.MiniPool) == MAX_POOL_SIZE
}

func (pool DefualtTxPool) GetPendingLastNonce(address string) uint64 {
	if len(pool.Pendings[address]) == 0 {
		return 0
	}
	lastBlk := pool.Pendings[address][len(pool.Pendings[address])-1]
	lastTx := lastBlk.Txs[len(lastBlk.Txs)-1]
	return lastTx.Nonce
}

func VerifyNonce(lastNonce uint64, nonce uint64) bool {
	return nonce == lastNonce+1
}

// PollFromQueue Put the transactions to pending
func (pool *DefualtTxPool) PollFromQueue(tx entity.Transaction, address string) {
	// Queue has been sorted by nonce
	nonce := tx.Nonce
	for i := 0; i < len(pool.Queues[address]); i++ {
		firstTx := pool.Queues[address][0]
		if VerifyNonce(pool.GetPendingLastNonce(address), nonce) {
			pool.PollFromQueueInsert(firstTx, address)
			pool.Queues[address] = append(pool.Queues[address][1:])
			nonce += 1
			if len(pool.Queues[address]) == 0 {
				pool.DeleteFromQueue(address)
			}
		}
	}
}

func (pool *DefualtTxPool) PollFromQueueInsert(tx entity.Transaction, address string) {
	rowBlks := pool.Pendings[address]
	for i := len(rowBlks) - 1; i >= 0; i-- {
		if rowBlks[i].GasPrice == tx.Gas && VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {
			rowBlks[i].Txs = append(rowBlks[i].Txs, tx)
		} else {
			pool.CreateBlk(tx, address)
		}
	}
}

// DeleteFromPending Delete the transaction from the pending list
func (pool *DefualtTxPool) DeleteFromPending(tx entity.Transaction) {
	address := typeconv.FromBytes[string](tx.From) //TODO
	rowBlks := pool.Pendings[address]
	for i := 0; i < len(rowBlks); i++ {
		if rowBlks[i].GasPrice == tx.Gas {
			for j := 0; j < len(rowBlks[i].Txs); j++ {
				if rowBlks[i].Txs[j].Nonce == tx.Nonce {
					rowBlks[i].Txs = append(rowBlks[i].Txs[:j], rowBlks[i].Txs[j+1:]...)
				}
			}
		}
	}
}

func (pool *DefualtTxPool) DeleteFromQueue(address string) {
	delete(pool.Queues, address)
}
