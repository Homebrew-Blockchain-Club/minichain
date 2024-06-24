package typeconv

import (
	"bytes"
	"encoding/gob"
)

// 将任意类型转换为[]byte
// 不要尝试用此函数转换公私钥
// 不要尝试用此函数转换big.Int
func ToBytes(data any) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(data)
	return buf.Bytes()
}

// 将[]byte转换为任意类型
// 不要尝试用此函数转换公私钥
// 不要尝试用此函数转换big.Int
func FromBytes[T any](data []byte) T {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var ret T
	dec.Decode(&ret)
	return ret
}

func ToHex(bytes []byte) []byte {
	result := make([]byte, len(bytes)*2)

	for i, b := range bytes {
		highNibble := (b >> 4) & 0x0F
		lowNibble := b & 0x0F
		result[2*i] = highNibble
		result[2*i+1] = lowNibble
	}

	return result
}

func FromHex(data []byte) []byte {

	result := make([]byte, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		highNibble := data[i] & 0x0F
		lowNibble := data[i+1] & 0x0F
		result[i/2] = (highNibble << 4) | lowNibble
	}

	return result
}
