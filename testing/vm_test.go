package testing

import (
	"fmt"
	"os"
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/hasher"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/Homebrew-Blockchain-Club/minichain/vm"
)

func TestVM(t *testing.T) {
	mpt := ds.NewMPT()
	code, err := os.ReadFile("test2.wasm")
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic("read file failed")
	}
	account := entity.Account{
		Code: code,
	}
	mpt.Update([]byte{1, 1, 4, 5, 1, 4}, typeconv.ToBytes(account))
	mpt.Commit()
	vm := vm.NewVM()
	vm.Run([]byte{1, 1, 4, 5, 1, 4}, hasher.Hash(typeconv.ToBytes(mpt)), 9, "Test", nil)
}
func TestVMOutOfGas(t *testing.T) {
	mpt := ds.NewMPT()
	code, err := os.ReadFile("test.wasm")
	if err != nil {
		fmt.Printf("%s", err.Error())
		panic("read file failed")
	}
	account := entity.Account{
		Code: code,
	}
	mpt.Update([]byte{1, 1, 4, 5, 1, 4}, typeconv.ToBytes(account))
	mpt.Commit()
	vm := vm.NewVM()
	vm.Run([]byte{1, 1, 4, 5, 1, 4}, hasher.Hash(typeconv.ToBytes(mpt)), 7, "Test", nil)
}
