package vm

import (
	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

// reuse code written by others
// func transferImpl(from, to []byte, amount uint64) error {
// 	accountByte := ds.GetMPT(VMContext.blk.Header.StateRoot).Query(from)
// 	txmpt := ds.GetMPT(VMContext.blk.Header.TransactionRoot)
// 	txmpt.Update()
// 	recvmpt := ds.GetMPT(VMContext.blk.Header.RecipientRoot)
// 	if accountByte == nil {
// 		return fmt.Errorf("account %s doesn't exist", string(from))
// 	}

//		account := typeconv.FromBytes[entity.Account](accountByte)
//		if account.Balance < amount {
//			return fmt.Errorf("account %s doesn't have enough balance", string(from))
//		}
//	}
func setImpl(key, val string) {
	account := typeconv.FromBytes[entity.Account](ds.GetMPT(VMContext.blk.Header.StateRoot).Query(VMContext.address))
	var stateroot *ds.MPT
	if tmp := ds.GetMPT(account.StorageRoot); tmp == nil {
		stateroot := ds.NewMPT()
	} else {
		stateroot = tmp
	}
}
