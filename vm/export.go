package vm

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L${SRCDIR}/lib -lminichain-wasm3 -lm3 -luvwasi_a -luv_a
#include "m3_env.h"
#include "export.h"
*/
import "C"
import "unsafe"

var vmContext struct {
	// setImpl      func([]byte, []byte)
	// getImpl      func([]byte) []byte
	// transferImpl func()
	// callImpl     func(address []byte, function string, argv []string) (string, error)
	vm      *VM
	address []byte
}

//export set
func set(key, val *C.char) {
	Key := C.GoString(key)
	//defer C.free(unsafe.Pointer(key))
	Val := C.GoString(val)
	//defer C.free(unsafe.Pointer(key))
	vmContext.vm.setImpl(vmContext.address, []byte(Key), []byte(Val))
	//vmContext.setImpl([]byte(Key), []byte(Val))
}

//export get
func get(key *C.char) *C.char {
	Key := C.GoString(key)
	//defer C.free(unsafe.Pointer(key))
	ret := vmContext.vm.getImpl(vmContext.address, []byte(Key))
	return C.CString(string(ret))
}

//export call
func call(address *C.char, function *C.char, argc C.int, argv **C.char) *C.char {
	addressgo := C.GoString(address)
	//defer C.free(unsafe.Pointer(address))
	functiongo := C.GoString(function)
	//defer C.free(unsafe.Pointer(function))
	//defer C.free(unsafe.Pointer(argv))
	argvgo := make([]string, 0)
	for i := 0; i < int(argc); i++ {
		argvgo = append(argvgo, C.GoString(*argv))
		defer C.free(unsafe.Pointer(*argv))
		//fmt.Printf("arg[%d]: %s\n", i, arg)
		argv = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(argv)) + unsafe.Sizeof(*argv)))
	}
	ret, err := vmContext.vm.Call([]byte(addressgo), functiongo, argvgo)
	if err == nil {
		return C.CString(ret)
	} else {
		return C.CString(err.Error())
		// switch ret.Error() {
		// case "no such account":
		// 	return 1
		// case "invalid bytecode":
		// 	return 2
		// case "out of gas":
		// 	return 3
		// }
	}
}
