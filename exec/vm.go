package exec

type VM struct {
}

func NewVM() VM {
	return VM{}
}

func (*VM) Run(code []byte, gaslimit uint64) []byte {
	return nil
}
