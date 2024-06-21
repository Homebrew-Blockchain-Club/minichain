package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
)

func TestMPT(t *testing.T) {
	tr := ds.NewMPT()
	tr.Update([]byte{1, 1, 4}, []byte("514"))
	tr.Update([]byte{1, 9, 1}, []byte("9810"))
	tr.Commit()
	println(tr.Proof())
}
