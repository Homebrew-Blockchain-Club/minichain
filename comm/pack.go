package comm

type Package struct {
	// 包的类型，是区块包还是交易包
	Type string
	//包的具体数据
	Data []byte
	//包所附带的公钥 用此公钥来解锁这个包
	//并且这个公钥同时也是发送方的地址
	Pubkey []byte
}
