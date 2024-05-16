// merkletree是数据结构的具体实现
package merkletree

import (
	"github.com/Homebrew-Blockchain-Club/minichain/norm/array"
	"github.com/Homebrew-Blockchain-Club/minichain/norm/calc"
	"github.com/Homebrew-Blockchain-Club/minichain/norm/hash"
	helperds "github.com/Homebrew-Blockchain-Club/minichain/norm/helper_ds"
	"github.com/Homebrew-Blockchain-Club/minichain/norm/kvstore"
	typeconvertor "github.com/Homebrew-Blockchain-Club/minichain/norm/type_convertor"
)

type MerkleTree interface {
	Root() []byte
	NewNode([]byte)
	Exist([]byte) bool
	DeleteNode([]byte)
	UpdateNode(old, new []byte)
	GetProof([]byte) [][]byte
}

type MerkleTreeImpl struct {
	db   kvstore.LevelDB
	size uint32
	root uint32
	delq *helperds.PriorityQueue[uint32]
}

// 使用指定的数据库目录，创建一颗空树
func NewMerkleTreeImpl(storage_path string) *MerkleTreeImpl {
	tr := &MerkleTreeImpl{
		db:   *kvstore.NewLevelDB(storage_path),
		size: 0,
		root: 0,
		delq: helperds.NewPriorityQueue(func(a, b uint32) bool {
			return a < b
		}),
	}
	tr.db.Put([]byte("size"), typeconvertor.Int2Byte(tr.size))
	tr.db.Put([]byte("root"), typeconvertor.Int2Byte(tr.size))
	return tr
}

// 从指定的数据库目录中反序列化出一棵树
func InitFromLevelDB(storage_path string) *MerkleTreeImpl {
	db := kvstore.NewLevelDB(storage_path)
	delq := helperds.NewPriorityQueue(func(a, b uint32) bool {
		return a < b
	})
	size, _ := db.Get([]byte("size"))
	root, _ := db.Get([]byte("root"))
	for i := 1; i <= int(typeconvertor.Byte2Int(size)); i++ {
		if has, _ := db.Has(typeconvertor.Int2Byte(uint32(i)*2 - 1)); !has {
			delq.Push(uint32(i)*2 - 1)
		}
	}
	tr := &MerkleTreeImpl{
		db:   *db,
		size: typeconvertor.Byte2Int(size),
		root: typeconvertor.Byte2Int(root),
		delq: delq,
	}
	return tr
}

// 传入内容 得到树中对应的节点哈希值
func (tr MerkleTreeImpl) GetNode(content []byte) []byte {
	key := hash.Sha3Slice256(content)
	pos, _ := tr.db.Get(key)
	key = array.ConcatHash(key, pos)
	return key
}

// 得到根节点的哈希值
func (tr MerkleTreeImpl) Root() []byte {
	// TODO
	pos, _ := tr.db.Get([]byte("root"))
	hash, _ := tr.db.Get(pos)
	return hash
}

// 增更删共用的子过程
// 自下而上的处理每个节点
// 填入节点在本层的下标和层级号 fa<<log则得到总下标
func (tr *MerkleTreeImpl) updateSubproc(fa uint32, log int) {
	//分别判断左、右节点是否存在
	hasl, _ := tr.db.Has(typeconvertor.Int2Byte(calc.Fa2Lson(fa) << (log - 1)))
	hasr, _ := tr.db.Has(typeconvertor.Int2Byte(calc.Fa2Rson(fa) << (log - 1)))
	if !(hasl || hasr) { //如果都不存在，说明此时是删除函数的调用，而本节点不存在任何儿子，因此删除本节点
		buf := typeconvertor.Int2Byte(fa << log)
		key, _ := tr.db.Get(buf)
		tr.db.Delete(buf)
		tr.db.Delete(key)
		key = array.ConcatHash(key, buf)
		tr.db.Delete(key)
	} else { //否则重新根据子节点的hash值，重新计算本节点的哈希值
		key := make([]byte, 0)
		if hasl {
			key, _ = tr.db.Get(typeconvertor.Int2Byte(calc.Fa2Lson(fa) << (log - 1)))
			if !hasr {
				key = array.CopyAppend(key, key...)
			}
		}
		if hasr {
			rkey, _ := tr.db.Get(typeconvertor.Int2Byte(calc.Fa2Rson(fa) << (log - 1)))
			key = array.CopyAppend(key, rkey...)
			if !hasl {
				key = array.CopyAppend(key, key...)
			}
		}
		key = hash.Sha3Slice256(key)
		buf := typeconvertor.Int2Byte(fa << log)
		if has, _ := tr.db.Has(buf); has { //如有需要将自动构造新的节点
			dkey, _ := tr.db.Get(buf)
			tr.db.Delete(dkey)
		}
		tr.db.Put(key, buf)
		tr.db.Put(buf, key)
	}
}

// 根据内容新建节点
func (tr *MerkleTreeImpl) NewNode(content []byte) {
	var pos uint32
	var buf []byte
	//首先检查删除队列中是否存在可用的节点 否则树容量扩展 并使用新节点
	if tr.delq.Empty() {
		tr.size++
		pos = tr.size*2 - 1
		tr.db.Put([]byte("size"), typeconvertor.Int2Byte(tr.size))
		tr.root = (1 << calc.Log2(tr.size)) // when size is full
		if tr.root < tr.size {
			tr.root <<= 1
		}
		tr.db.Put([]byte("root"), typeconvertor.Int2Byte(tr.root))
	} else {
		pos = tr.delq.Pop().(uint32)
	}
	buf = typeconvertor.Int2Byte(pos)
	key := hash.Sha3Slice256(content)
	tr.db.Put(key, buf)
	key = array.ConcatHash(key, buf)
	tr.db.Put(buf, key)
	tr.db.Put(key, content)
	cur := pos
	log := 0
	for cur<<log != tr.root { //自下而上逐级递归
		log++
		if calc.IsLson(cur) {
			cur = calc.Lson2Fa(cur)
		} else {
			cur = calc.Rson2Fa(cur)
		}
		tr.updateSubproc(cur, log)
	}
}

// 检查树中是否存在对应的节点
func (tr MerkleTreeImpl) Exist(content []byte) bool {
	key := hash.Sha3Slice256(content)
	has, _ := tr.db.Has(key)
	return has
}

// 根据内容删除节点
func (tr *MerkleTreeImpl) DeleteNode(content []byte) {
	key := hash.Sha3Slice256(content)
	pos, _ := tr.db.Get(key)
	tr.delq.Push(typeconvertor.Byte2Int(pos))
	tr.db.Delete(key)
	tr.db.Delete(pos)
	key = array.ConcatHash(key, pos)
	tr.db.Delete(key)
	cur := typeconvertor.Byte2Int(pos)
	log := 0
	for cur<<log != tr.root {
		log++
		if calc.IsLson(cur) {
			cur = calc.Lson2Fa(cur)
		} else {
			cur = calc.Rson2Fa(cur)
		}
		tr.updateSubproc(cur, log)
	}
	log = int(calc.Lowcnt(tr.root))
	cur = tr.root >> log
	//如有需要则重新计算根节点 从原来的根节点逐步下降 删除只有一个孩子的节点 直到当前节点有两个孩子
	for {
		hasl, _ := tr.db.Has(typeconvertor.Int2Byte(calc.Fa2Lson(cur) << (log - 1)))
		hasr, _ := tr.db.Has(typeconvertor.Int2Byte(calc.Fa2Rson(cur) << (log - 1)))
		if hasl && hasr {
			break
		}
		buf := typeconvertor.Int2Byte(cur << log)
		key, _ := tr.db.Get(buf)
		tr.db.Delete(key)
		tr.db.Delete(buf)
		key = array.ConcatHash(key, buf)
		tr.db.Delete(key)
		if hasl {
			cur = calc.Fa2Lson(cur)
		} else {
			cur = calc.Fa2Rson(cur)
		}
		log--
		if log == 0 { //不删除叶子结点
			break
		}
		tr.root = cur << log
	}

}

// 更新节点内容
func (tr *MerkleTreeImpl) UpdateNode(old, new []byte) {
	key := hash.Sha3Slice256(old)
	pos, _ := tr.db.Get(key)
	tr.db.Delete(key)
	key = array.ConcatHash(key, pos)
	tr.db.Delete(key)
	cur := typeconvertor.Byte2Int(pos)
	key = hash.Sha3Slice256(new)
	tr.db.Put(key, pos)
	key = array.ConcatHash(key, pos)
	tr.db.Put(pos, key)
	tr.db.Put(key, new)
	log := 0
	for cur<<log != tr.root {
		log++
		if calc.IsLson(cur) {
			cur = calc.Lson2Fa(cur)
		} else {
			cur = calc.Rson2Fa(cur)
		}
		tr.updateSubproc(cur, log)
	}
}

// 求取节点的证明以证明根节点正确性
func (tr MerkleTreeImpl) GetProof(content []byte) [][]byte {
	var ret [][]byte
	key := hash.Sha3Slice256(content)
	pos, _ := tr.db.Get(key)
	cur := typeconvertor.Byte2Int(pos)
	log := 0
	subproc := func(cur, sib uint32) {
		buf := typeconvertor.Int2Byte(sib)
		if has, _ := tr.db.Has(buf); has {
			sibkey, _ := tr.db.Get(buf)
			ret = array.CopyAppend(ret, sibkey)
		} else {
			buf = typeconvertor.Int2Byte(cur)
			key, _ := tr.db.Get(buf)
			ret = array.CopyAppend(ret, key)
		}
	}
	for cur<<log != tr.root {
		if calc.IsLson(cur) {
			subproc(cur<<log, calc.Fa2Rson(calc.Lson2Fa(cur))<<log)
			cur = calc.Lson2Fa(cur)
		} else {
			subproc(cur<<log, calc.Fa2Lson(calc.Rson2Fa(cur))<<log)
			cur = calc.Rson2Fa(cur)
		}
		log++
	}
	return ret
}
