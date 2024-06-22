package hasher

import "github.com/ethereum/go-ethereum/crypto"

// 对数据进行sha3哈希 返回哈希的结果
func Hash(data []byte) []byte {
	return crypto.Keccak256(data)
}
