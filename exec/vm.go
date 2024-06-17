package exec

type VM struct {
}

// 创建一个新的解释器
func newVM() VM {
	return VM{}
}

// TODO 区块链上下文如何引入
// 用此解释器执行VM字节码，并给出执行的合约账户地址和gas上限
func (*VM) run(account []byte, code []byte, gaslimit uint64, stateroot []byte) []byte {
	return nil
}
