package exec

import (
	_ "encoding/binary"
	"math/big"
)

// check 检查哈希是否满足给定的难度约束
func check(hash []byte, difficulty uint64) bool {
	// 计算目标值
	target := new(big.Int).Div(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(int64(difficulty)))

	// 将哈希值转换为大整数
	hashInt := new(big.Int).SetBytes(hash)

	// 比较哈希值是否小于目标值
	return hashInt.Cmp(target) == -1
}
