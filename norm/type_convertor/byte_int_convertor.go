// typeconvertor提供了类型转换实用函数
package typeconvertor

import "encoding/binary"

func Byte2Int(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data[:4])
}
func Int2Byte(data uint32) []byte {
	ret := make([]byte, 4)
	binary.LittleEndian.PutUint32(ret, data)
	return ret
}
