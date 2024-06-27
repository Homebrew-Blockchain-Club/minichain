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
func (vm *VM) getStorageRoot(address []byte) *ds.MPT {
	var rt *ds.MPT
	account := typeconv.FromBytes[entity.Account](vm.stateRoot.Query(address))
	if root, ok := vm.pendingStorage[string(address)]; ok {
		rt = ds.GetMPT(root)
	} else if ds.GetMPT(account.StorageRoot) != nil {
		rt = ds.NewMPTFromPrevious(account.StorageRoot)
	} else {
		rt = ds.NewMPT()
	}
	return rt
}
func (vm *VM) setImpl(address, key, val []byte) {
	rt := vm.getStorageRoot(address)
	rt.Update(key, val)
	vm.pendingStorage[string(address)] = typeconv.ToBytes(*rt)
}

func (vm *VM) getImpl(address, key []byte) []byte {
	return vm.getStorageRoot(address).Query(key)
}
