// merkletree_test是测试包
package merkletree_test

import (
	"crypto/rand"
	"testing"

	merkletree "github.com/Homebrew-Blockchain-Club/minichain/norm"
	"github.com/Homebrew-Blockchain-Club/minichain/norm/array"
)

func generateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	_, _ = rand.Read(bytes) // 生成随机字节
	// if err != nil {
	// 	return nil, err
	// }
	return bytes
}
func TestMerkleTree(t *testing.T) {
	tr := merkletree.NewMerkleTreeImpl("testdb")
	size := 10000
	sel := generateRandomBytes(100)
	tr.NewNode(sel)
	del := generateRandomBytes(100)
	tr.NewNode(del)
	key := tr.GetNode(sel)
	proof := tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: new node small mount")

		}
	}
	var data [][]byte
	for i := 2; i < size; i++ {
		x := generateRandomBytes(100)
		tr.NewNode(x)
		data = append(data, x)
	}
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: new node")

		}
	}
	old := sel
	sel = generateRandomBytes(100)
	tr.UpdateNode(old, sel)
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: update node")

		}
	}
	tr.DeleteNode(del)
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: delete node")

		}
	}
	tr.NewNode(generateRandomBytes(100))
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: insert node")

		}
	}
	for i := 5; i <= 5000; i++ {
		tr.DeleteNode(data[i])
	}
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: multiple delete node")

		}
	}
	for i := 5001; i <= 7000; i++ {
		tr.UpdateNode(data[i], generateRandomBytes(100))
	}
	key = tr.GetNode(sel)
	proof = tr.GetProof(sel)
	for _, ele := range proof {
		key = array.ConcatHash(key, ele)
	}
	for k, ele := range tr.Root() {
		if key[k] != ele {
			panic("something wrong: multiple update node")

		}
	}
	for i := 5; i <= 7000; i++ {
		if tr.Exist(data[i]) {
			panic("something wrong: something is still there")
		}
	}

	println("all done")
}
