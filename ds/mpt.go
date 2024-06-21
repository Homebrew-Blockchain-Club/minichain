package ds

import (
	"bytes"

	"github.com/Homebrew-Blockchain-Club/minichain/hasher"
	storage "github.com/Homebrew-Blockchain-Club/minichain/store"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

type node interface {
	isDirty() bool
	setDirty()
}

// 将[]byte转为string后取得
var nodeCache map[string]node = make(map[string]node)

// 被保存了的节点必然为被哈希的节点
type (
	fullNode struct {
		Children [17]string //子节点哈希
		Dirty    bool
	}
	shortNode struct {
		Key   []byte
		Val   string //子节点哈希
		Dirty bool
	}
	valueNode struct {
		Val   []byte
		Dirty bool
	}
)

func (x *fullNode) isDirty() bool {
	return x.Dirty
}
func (x *fullNode) setDirty() {
	x.Dirty = true
}
func (x *fullNode) copy() *fullNode {
	copy := *x
	copy.Dirty = false
	return &copy
}
func (x *shortNode) isDirty() bool {
	return x.Dirty
}
func (x *shortNode) setDirty() {
	x.Dirty = true
}
func (x *valueNode) isDirty() bool {
	return x.Dirty
}
func (x *valueNode) setDirty() {
	x.Dirty = true
}

type MPT struct {
	Root      string //root的哈希值
	Committed bool
}

func getNode(hash string) node {
	ret := nodeCache[hash]
	if ret != nil {
		return ret
	}
	if nodebyte := storage.Query(typeconv.ToBytes(hash)); nodebyte != nil {
		typebyte := storage.Query(typeconv.ToBytes(hash + "type"))
		switch string(typebyte) {
		case "F":
			cur := typeconv.FromBytes[fullNode](nodebyte)
			nodeCache[hash] = &cur
		case "S":
			cur := typeconv.FromBytes[shortNode](nodebyte)
			nodeCache[hash] = &cur
		case "V":
			cur := typeconv.FromBytes[valueNode](nodebyte)
			nodeCache[hash] = &cur
		}
		return nodeCache[hash]
	} else {
		return nil
	}
}
func hash(x node, key string) string {
	switch x := x.(type) {
	case *shortNode:
		return string(hasher.Hash([]byte(key + x.Val)))
	case *fullNode:
		str := key
		for _, ele := range x.Children {
			str += ele
		}
		return string(hasher.Hash([]byte(str)))
	case *valueNode:
		return string(hasher.Hash([]byte(key + string(x.Val))))
	}
	return ""
}

// 持久保存节点需要另外实现
func setNode(hash string, cur node) {
	nodeCache[hash] = cur
}
func GetMPT(hash []byte) *MPT {
	ret := typeconv.FromBytes[MPT](storage.Query(hash))
	return &ret
}

// 创建新的MPT 用于交易树和收据树
func NewMPT() *MPT {
	return &MPT{
		Root:      "",
		Committed: false,
	}
}
func (tr *MPT) IsComitted() bool {
	return tr.Committed
}

// 创建新的MPT，用于状态树
func NewStateMPT(top *Block) *MPT {
	return &MPT{
		Root:      typeconv.FromBytes[string](top.Header.StateRoot),
		Committed: false,
	}
}
func (tr *MPT) commit(cur node, key []byte) {
	if cur.isDirty() {
		return
	}
	switch cur := cur.(type) {
	case nil:
		return
	case *fullNode:
		storage.Store([]byte(hash(cur, string(key))), typeconv.ToBytes(cur))
		storage.Store([]byte(hash(cur, string(key)+"type")), []byte("F"))
		for s, nxt := range cur.Children {
			tr.commit(getNode(nxt), append(key, byte(s)))
		}
	case *shortNode:
		storage.Store([]byte(hash(cur, string(key))), typeconv.ToBytes(cur))
		storage.Store([]byte(hash(cur, string(key)+"type")), []byte("S"))
		tr.commit(getNode(cur.Val), append(key, cur.Key...))
	case *valueNode:
		storage.Store([]byte(hash(cur, string(key))), typeconv.ToBytes(cur))
		storage.Store([]byte(hash(cur, string(key)+"type")), []byte("V"))
	}
}

// 将本树持久化
func (tr *MPT) Commit() {}
func prefixLen(a, b []byte) int {
	i, len := 0, min(len(a), len(b))
	for ; i < len; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}
func (tr *MPT) update(cur node, prefix, key []byte, val node) (bool, node) {
	if len(key) == 0 {
		if bytes.Equal(val.(*valueNode).Val, cur.(*valueNode).Val) {
			//cur.(*valueNode).Dirty = true

			return false, cur
		} else {
			cur = val
			setNode(hash(cur, string(prefix)), cur)
			return true, cur
		}
	}
	if cur != nil {
		cur.setDirty()
	}
	switch cur := cur.(type) {
	case *shortNode:
		matchlen := prefixLen(key, cur.Key)
		if matchlen == len(cur.Key) {
			// if son := getNode(cur.Val); son != nil {
			// 	son.setDirty()
			// }
			nprefix := append(prefix, key[:matchlen]...)
			dirty, nxt := tr.update(getNode(cur.Val), nprefix, key[matchlen:], val)
			if !dirty {
				return false, cur
			} else {
				new := &shortNode{
					Key:   cur.Key,
					Val:   hash(nxt, string(nprefix)),
					Dirty: false,
				}
				setNode(hash(new, string(prefix)), new)
				return true, new
			}
		}
		branch := &fullNode{Dirty: false}
		{
			nprefix := append(prefix, cur.Key[:matchlen+1]...)
			_, child := tr.update(nil, nprefix, cur.Key[matchlen+1:], getNode(cur.Val))
			branch.Children[cur.Key[matchlen]] = hash(child, string(nprefix))
		}
		{
			nprefix := append(prefix, key[:matchlen+1]...)
			_, child := tr.update(nil, nprefix, key[matchlen+1:], val)
			branch.Children[key[matchlen]] = hash(child, string(nprefix))
		}
		branchHash := hash(branch, string(append(prefix, key[:matchlen]...)))
		setNode(branchHash, branch)
		if matchlen == 0 {
			return true, branch
		}
		link := &shortNode{key[:matchlen], branchHash, false}
		setNode(hash(link, string(prefix)), link)
		return true, link
	case *fullNode:
		nprefix := append(prefix, key[0])
		// if son := getNode(string(nprefix)); son != nil {
		// 	son.setDirty()
		// }
		dirty, nxt := tr.update(getNode(cur.Children[key[0]]), nprefix, key[1:], val)
		if !dirty {
			return false, cur
		} else {
			cur = cur.copy()
			cur.Children[key[0]] = hash(nxt, string(nprefix))
			setNode(hash(cur, string(prefix)), cur)
			return true, cur
		}
	case nil:
		nprefix := append(prefix, key...)
		valhash := hash(val, string(nprefix))
		setNode(valhash, val)
		cur = &shortNode{key, valhash, false}
		setNode(hash(cur, string(prefix)), cur)
		return true, cur
	}
	return false, nil
	//return nil
}

// 向对应的MPT存储键值对。若不存在则会创建新的，若存在则会覆盖原数据
func (tr *MPT) Update(key, val []byte) {
	if tr.Committed {
		panic("attempt to change a trie that has been already committed")
	}
	root := getNode(tr.Root)
	Val := &valueNode{
		Val:   val,
		Dirty: false,
	}
	_, root = tr.update(root, []byte{}, key, Val)
	newHash := hash(root, "")
	tr.Root = newHash
}
func (tr *MPT) query(cur node, key []byte) []byte {
	switch cur := cur.(type) {
	case nil:
		return nil
	case *valueNode:
		return cur.Val
	case *shortNode:
		if len(key) < len(cur.Key) || !bytes.Equal(cur.Key, key[:len(cur.Key)]) {
			return nil // not found
		}
		return tr.query(getNode(cur.Val), key[len(cur.Key):])
	case *fullNode:
		return tr.query(getNode(cur.Children[key[0]]), key[1:])
	}
	return nil
}

// 给出key查询节点的信息，若不存在则返回nil
func (tr *MPT) Query(key []byte) []byte {
	return tr.query(getNode(tr.Root), key)
}
func (tr *MPT) proof(cur node, prefix []byte) bool {
	switch cur := cur.(type) {
	case *valueNode:
		return getNode(hash(cur, string(prefix))) == cur
	case *fullNode:
		for s, ele := range cur.Children {
			if ele == "" {
				continue
			}
			if !tr.proof(getNode(ele), append(prefix, byte(s))) {
				return false
			}
		}
		return getNode(hash(cur, string(prefix))) == cur
	case *shortNode:
		if !tr.proof(getNode(cur.Val), append(prefix, cur.Key...)) {
			return false
		}
		return getNode(hash(cur, string(prefix))) == cur
	}
	return false
}

// 验证这个MPT是否可信
func (tr MPT) Proof() bool {
	return tr.proof(getNode(tr.Root), []byte{})
}
