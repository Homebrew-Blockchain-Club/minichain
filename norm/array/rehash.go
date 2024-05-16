package array

import "github.com/Homebrew-Blockchain-Club/minichain/norm/hash"

func ConcatHash(a, b []byte) []byte {
	a = CopyAppend(a, b...)
	return hash.Sha3Slice256(a)
}
