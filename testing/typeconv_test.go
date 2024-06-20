package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

func TestTypeConv(t *testing.T) {
	x := 5
	y := typeconv.ToBytes(5)
	x = typeconv.FromBytes[int](y)
	println(x)
}
