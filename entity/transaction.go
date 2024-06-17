package entity

type Transaction struct {
	Nonce    uint64
	GasLimit uint64
	To       []byte
	Data     []byte
	//没有d r s 因为通信的时候将使用另一个结构体进行包装
}
