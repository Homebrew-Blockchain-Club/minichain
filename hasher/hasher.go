package hasher

import "golang.org/x/crypto/sha3"

// 对数据进行sha3哈希 返回哈希的结果
func Hash(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}
