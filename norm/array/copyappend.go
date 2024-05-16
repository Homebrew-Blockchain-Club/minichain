// array提供切片的实用函数
package array

// 将切片a与切片b连接起来，返回新切片，不改变a的值
func CopyAppend[T any](arr []T, x ...T) []T {
	ret := make([]T, 0)
	ret = append(ret, arr...)
	ret = append(ret, x...)
	return ret
}
