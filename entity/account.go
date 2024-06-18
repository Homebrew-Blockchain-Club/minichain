package entity

type Account struct {
	Nonce       uint64
	Balance     uint64
	StorageRoot []byte
	//如果账户并非合约账户 则Code==nil
	Code []byte
}
