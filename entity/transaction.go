package entity

type Transaction struct {
	Nonce    uint64
	GasLimit uint64
	From     []byte
	To       []byte
	Data     []byte
	Sign     []byte
}
