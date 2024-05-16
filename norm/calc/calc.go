// calc提供了默克尔树进行节点遍历所需要的数学函数
package calc

func Log2(x uint32) uint32 {
	var ret uint32 = 0
	for x > 1 {
		x /= 2
		ret++
	}
	return ret
}
func lowbit(x uint32) uint32 {
	return uint32(int32(x) & int32(-x))
}

// 计算x属于哪一层
func Lowcnt(x uint32) uint32 {
	return Log2(lowbit(x))
}

// 计算x作为左子节点，其父节点y的值
func Lson2Fa(x uint32) uint32 {
	return (x + 1) / 2
}

// 计算x作为右子节点，其父节点y的值
func Rson2Fa(x uint32) uint32 {
	return (x - 1) / 2
}

// 计算x作为父节点，其左子结点的值
func Fa2Lson(x uint32) uint32 {
	return x*2 - 1
}

// 计算x作为父节点，其右子节点的值
func Fa2Rson(x uint32) uint32 {
	return x*2 + 1
}

// 判断x是否为某个节点的左子节点
func IsLson(x uint32) bool {
	return (x+1)/2%2 == 1
}
