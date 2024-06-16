package hasher

import "golang.org/x/crypto/sha3"

func Hash(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}
