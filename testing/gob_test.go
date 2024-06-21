package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
)

func TestGob(t *testing.T) {
	a := typeconv.ToBytes(1)
	println(a)
}
