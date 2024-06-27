package testing

import (
	"fmt"
	"os"
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/Homebrew-Blockchain-Club/minichain/vm"
)

func TestVM(t *testing.T) {
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

	vm := vm.NewVM([]byte{}, mpt, 0, 100)

	str, err := vm.Call([]byte{1, 1, 4, 5, 1, 4}, "Test", nil)
	if err == nil {
		fmt.Println(str)
	}
}
func TestVMSet(t *testing.T) {
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
	//mpt.Commit()

	vm := vm.NewVM([]byte{}, mpt, 0, 1000)

	_, err = vm.Call([]byte{1, 1, 4, 5, 1, 4}, "Test", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	vm.Finish()
	st := typeconv.FromBytes[entity.Account](mpt.Query([]byte{1, 1, 4, 5, 1, 4})).StorageRoot
	if st == nil {
		t.Error("stateroot unchanged")
		return
	}
	fmt.Println(string(ds.GetMPT(st).Query([]byte("114"))))
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
	vm := vm.NewVM([]byte{}, mpt, 0, 100)
	vm.Call([]byte{1, 1, 4, 5, 1, 4}, "Test", nil)
}
