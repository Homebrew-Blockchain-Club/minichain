package ds

type node interface{}
type MPT struct {
	root *node
}

// 创建新的MPT 用于交易树和收据树
func NewMPT() *MPT {
	return &MPT{
		root: nil,
	}
}

// 创建新的MPT，用于状态树
func NewStateMPT(*Block) *MPT { return nil }

// 将本树持久化
func (tr *MPT) Commit() {}

// 向对应的MPT存储键值对。若不存在则会创建新的，若存在则会覆盖原数据
func (tr *MPT) Update(key, val []byte) {

}

// 给出key查询节点的信息，若不存在则返回nil
func (tr *MPT) Query([]byte) []byte {
	return nil
}

// 验证这个MPT是否可信
func (tr MPT) Proof() bool {
	return true
}
