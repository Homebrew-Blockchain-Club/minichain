package vm

type AbstractVM interface {
	Run(account []byte, code []byte, gaslimit uint64, stateroot []byte) []byte
}
