package vm

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L${SRCDIR}/lib -lminichain-wasm3 -lm3 -luvwasi_a -luv_a
#include "m3_env.h"
#include "export.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

type VM struct {
	env           *C.struct_M3Environment
	stateRoot     *ds.MPT
	gas, gaslimit uint64

	pendingStorage map[string][]byte //key 是待提交的账户StorageRoot
}

// 创建一个新的解释器
func NewVM(address []byte, stateRoot *ds.MPT, gas, gaslimit uint64) *VM {
	return &VM{
		env:            C.m3_NewEnvironment(),
		pendingStorage: make(map[string][]byte),
		stateRoot:      stateRoot,
		gas:            gas,
		gaslimit:       gaslimit,
	}
}
func (vm *VM) Finish() {
	//st := ds.GetMPT(vm.blk.Header.StateRoot)
	for key, val := range vm.pendingStorage {
		account := typeconv.FromBytes[entity.Account](vm.stateRoot.Query([]byte(key)))
		tr := typeconv.FromBytes[ds.MPT](val)
		trByte := tr.Commit()
		account.StorageRoot = trByte
		vm.stateRoot.Update([]byte(key), typeconv.ToBytes(account))
	}
	vm.Destruct()
}
func (vm *VM) Destruct() {
	C.m3_FreeEnvironment(vm.env)
}
func mkuintptr(arr [8]byte) uintptr {
	var ans uintptr = 0
	for i := 7; i >= 0; i-- {
		ans <<= 8
		ans += uintptr(arr[i])
	}
	//fmt.Printf("%d", int64(ans))
	return ans
}

// TODO 区块链上下文如何引入
// 用此解释器执行VM字节码，并给出执行的合约账户地址和gas上限
func (vm *VM) Call(address []byte, function string, argv []string) (string, error) {
	// prevSetImpl := vmContext.setImpl
	// vmContext.setImpl =vm.setImpl(address, key, val)

	// //C.set = (*[0]byte)(unsafe.Pointer(setWrapper))
	// defer func() { vmContext.setImpl = prevSetImpl }()
	// prevGetImpl := vmContext.getImpl
	// vmContext.getImpl =
	// 	vm.getImpl(address, key)

	// //C.get = (*[0]byte)(unsafe.Pointer(&vmContext.getImpl))
	// defer func() { vmContext.getImpl = prevGetImpl }()
	// prevCallImpl := vmContext.callImpl
	// vmContext.callImpl =vm.Call(address, function, argv)

	// //C.call = (*[0]byte)(unsafe.Pointer(&vmContext.callImpl))

	// defer func() { vmContext.callImpl = prevCallImpl }()
	preVM := vmContext.vm
	vmContext.vm = vm
	defer func() { vmContext.vm = preVM }()
	preAddress := vmContext.address
	vmContext.address = address
	defer func() { vmContext.address = preAddress }()
	cntch := make(chan uint64)
	runch := make(chan C.struct_ResultString)
	account := typeconv.FromBytes[entity.Account](vm.stateRoot.Query(address))
	if account.Code == nil {
		return "", fmt.Errorf("no such account")
	}
	modr := C.GetModuleFromBytecode(vm.env, (*C.uchar)(unsafe.Pointer(&account.Code[0])), C.size_t(len(account.Code)))
	if modr.state == C.ERROR {
		return "", fmt.Errorf("invalid bytecode")

	}
	functioncstr := C.CString(function)
	defer C.free(unsafe.Pointer(functioncstr))
	argvcstr := make([]*C.char, len(argv))
	for i, s := range argv {
		cstr := C.CString(s)
		defer C.free(unsafe.Pointer(cstr))
		argvcstr[i] = cstr
	}
	mod := (*C.struct_M3Module)(unsafe.Pointer(mkuintptr(modr.data)))
	rt := C.m3_NewRuntime(vm.env, 64*1024, nil)
	defer C.m3_FreeRuntime(rt)
	ret := C.AttachToRuntime(mod, rt)
	if ret.state == C.ERROR {
		return "", fmt.Errorf(C.GoString((*C.char)(unsafe.Pointer(mkuintptr(ret.data)))))
	}
	go func() {
		for uint64(C.gascnt) <= vm.gaslimit {
		}
		cntch <- uint64(C.gascnt)
	}()
	go func() {
		if len(argvcstr) > 0 {
			runch <- C.RunFunction(mod, rt, vm.env, functioncstr, C.int(len(argv)), &argvcstr[0])
		} else {
			runch <- C.RunFunction(mod, rt, vm.env, functioncstr, C.int(len(argv)), nil)
		}
	}()
	select {
	case <-cntch:
		//fmt.Printf("out of gas\n")
		C.forceexit = 1
		//ret := <-runch
		//fmt.Printf("%s\n", C.GoString((*C.char)(unsafe.Pointer(mkuintptr(ret.data)))))
		return "", fmt.Errorf("out of gas")

	case ret := <-runch:
		if uint64(C.gascnt) > vm.gaslimit {
			//fmt.Printf("out of gas\n")
			return "", fmt.Errorf("out of gas")
		}
		if ret.state == C.SUCCESS {
			return C.GoString((*C.char)(unsafe.Pointer(mkuintptr(ret.data)))), nil
		} else {
			return "", fmt.Errorf(C.GoString((*C.char)(unsafe.Pointer(mkuintptr(ret.data)))))
		}
	}

}
