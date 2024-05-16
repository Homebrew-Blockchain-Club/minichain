// hash提供了求数据哈希值的实用函数
package hash

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

type HashPool interface {
	Get() hash.Hash
}

type DefaultHash struct{}

func (pool DefaultHash) Get() hash.Hash {
	return sha3.New256()
}

func Sha3Sum256(data []byte) []byte {
	h := sha3.New256()
	return h.Sum(data)
}

func Sha3Slice256(data []byte) []byte {
	h := sha3.New256() // 创建新的 SHA-256 哈希对象
	h.Write(data)      // 写入数据到哈希
	return h.Sum(nil)  // 获取哈希摘要
}
