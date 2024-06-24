package vm

type AbstractVM interface {
	Run(address []byte, stateroot []byte, gaslimit uint64, function string, argv []string)
}
