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
	address  string
}

type TxsBlks []SortedTxsBlk

func (td TxsBlks) Len() int {
	return len(td)
}
func (td TxsBlks) Less(i, j int) bool {
	f := true
	if td[i].address == td[j].address {
		f = td[i].Txs[len(td[i].Txs)-1].Nonce < td[j].Txs[0].Nonce
	}
	return td[i].GasPrice > td[j].GasPrice && f

}
func (td TxsBlks) Swap(i, j int) {
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
	Pendings map[string]TxsBlks // map to store pending transactions grouped by nonce
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
	if pool.Pendings == nil {
		pool.Pendings = make(map[string]TxsBlks)
	}

	// Trying to insert the transaction into the pending list
	if tx.Nonce == account.Nonce+1 || VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {

		if len(rowBlks) == 0 {
			pool.CreateTxsBlk(tx, address)
		} else {
			for i := len(rowBlks) - 1; i >= len(rowBlks)-1; i-- {
				if rowBlks[i].GasPrice == tx.Gas && VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {
					pool.Pendings[address][i].Txs = append(pool.Pendings[address][i].Txs, tx)
				} else {
					pool.CreateTxsBlk(tx, address)
				}
			}
		}
		pool.PollFromQueueToPending(address)

	} else if tx.Nonce > pool.GetPendingLastNonce(address)+1 {
		// If the nonce is out of order, add it to the queue
		queueTxs := pool.Queues[address]
		if len(queueTxs) == 0 {
			pool.Queues = make(map[string]Txs)
			queueTxs = pool.Queues[address]
		}
		pool.Queues[address] = append(pool.Queues[address], tx)
		sort.Sort(queueTxs)
	} else if tx.Nonce <= pool.GetPendingLastNonce(address) {
		// If the nonce is less than the last nonce, ignore it
		return
	}
	pool.AllSortedByGas()
}

// CreateTxsBlk Create Blk and append the first new transaction
func (pool *DefualtTxPool) CreateTxsBlk(tx entity.Transaction, address string) {
	newBlk := SortedTxsBlk{
		Txs:      make([]entity.Transaction, 0),
		GasPrice: tx.Gas,
		address:  address,
	}
	newBlk.Txs = append(newBlk.Txs, tx)
	pool.Pendings[address] = append(pool.Pendings[address], newBlk)

}

// Poll 从合约池中取出第一个交易，并从合约池中弹出该交易
func (pool *DefualtTxPool) Poll() entity.Transaction {
	tx := pool.MiniPool[0]

	// Delete the transaction from the pending list
	pool.DeleteFromPending(tx)

	// Update the miniPool
	pool.AllSortedByGas()
	return tx
}

func (pool *DefualtTxPool) Length() int {
	return len(pool.MiniPool)
}

func (pool *DefualtTxPool) AllSortedByGas() {
	allBlk := make(TxsBlks, 0)
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
	pool.MiniPool = make(Txs, 0)
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

// PollFromQueueToPending Put the transactions to pending
func (pool *DefualtTxPool) PollFromQueueToPending(address string) {
	// Queue has been sorted by nonce
	length := len(pool.Queues[address])
	for i := 0; i < length; i++ {
		firstTx := pool.Queues[address][0]
		if VerifyNonce(pool.GetPendingLastNonce(address), firstTx.Nonce) {
			pool.QueueInsertToPending(firstTx, address)
			pool.Queues[address] = append(pool.Queues[address][1:])

			if len(pool.Queues[address]) == 0 {
				pool.DeleteQueueRow(address)
			}
		}
	}
}

func (pool *DefualtTxPool) QueueInsertToPending(tx entity.Transaction, address string) {
	rowBlks := pool.Pendings[address]
	for i := len(rowBlks) - 1; i >= len(rowBlks)-1; i-- {
		if rowBlks[i].GasPrice == tx.Gas && VerifyNonce(pool.GetPendingLastNonce(address), tx.Nonce) {
			pool.Pendings[address][i].Txs = append(pool.Pendings[address][i].Txs, tx)
		} else {
			pool.CreateTxsBlk(tx, address)
		}
	}
}

// DeleteFromPending Delete the transaction from the pending list
func (pool *DefualtTxPool) DeleteFromPending(tx entity.Transaction) {
	address := common.BytesToAddress(tx.From).Hex()
	rowBlks := pool.Pendings[address]

	i := 0
	lengthTxs := len(rowBlks[i].Txs)
	if rowBlks[i].GasPrice == tx.Gas {
		for j := 0; j < lengthTxs; j++ {
			if rowBlks[i].Txs[j].Nonce == tx.Nonce {
				pool.Pendings[address][i].Txs = append(pool.Pendings[address][i].Txs[:j], pool.Pendings[address][i].Txs[j+1:]...)
				break
			}
		}
	}
	// Delete the Blk if it is empty
	if len(pool.Pendings[address][i].Txs) == 0 {
		pool.Pendings[address] = append(pool.Pendings[address][:i], pool.Pendings[address][i+1:]...)
	}

	// Delete the row if it is empty
	if len(pool.Pendings[address]) == 0 {
		delete(pool.Pendings, address)
	}
}

func (pool *DefualtTxPool) DeleteQueueRow(address string) {
	delete(pool.Queues, address)
}
