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
	env *C.struct_M3Environment
}

// 创建一个新的解释器
func NewVM() *VM {
	env := C.m3_NewEnvironment()
	return &VM{
		env: env,
	}
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
func (vm *VM) Run(address []byte, stateroot []byte, gaslimit uint64, function string, argv []string) {
	cntch := make(chan uint64)
	runch := make(chan C.struct_ResultString)
	account := typeconv.FromBytes[entity.Account](ds.GetMPT(stateroot).Query(address))
	if account.Code == nil {
		return
	}
	modr := C.GetModuleFromBytecode(vm.env, (*C.uchar)(unsafe.Pointer(&account.Code[0])), C.ulong(len(account.Code)))
	if modr.state == C.ERROR {
		return
	}
	functioncstr := C.CString(function)
	//defer C.free(unsafe.Pointer(&functioncstr))
	argvcstr := make([]*C.char, len(argv))
	for i, s := range argv {
		cstr := C.CString(s)
		//defer C.free(unsafe.Pointer(&cstr))
		argvcstr[i] = cstr
	}
	mod := (*C.struct_M3Module)(unsafe.Pointer(mkuintptr(modr.data)))
	rt := C.m3_NewRuntime(vm.env, 64*1024, nil)
	C.AttachToRuntime(mod, rt)
	go func() {
		for uint64(C.gascnt) <= gaslimit {

		}
		//println(uint64(C.gascnt))
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
		fmt.Printf("out of gas")
	case ret := <-runch:
		if ret.state == C.SUCCESS {
			tot := <-cntch
			println(tot)
			fmt.Printf("%s", C.GoString((*C.char)(unsafe.Pointer(mkuintptr(ret.data)))))
		}
	}
}
