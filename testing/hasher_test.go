package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/hasher"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

func TestHasher(t *testing.T) {
	a := hasher.Hash(typeconv.ToBytes(1))
	println(a)
}
