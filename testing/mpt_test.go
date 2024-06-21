package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
)

func TestMPT(t *testing.T) {
	tr := ds.NewMPT()
	tr.Update([]byte{1, 1, 4}, []byte("514"))
	tr.Update([]byte{1, 9, 1}, []byte("9810"))
	tr.Update([]byte{1, 2, 3}, []byte("测试你的码"))
	tr.Update([]byte{8, 1, 0}, []byte("丁真电子烟"))
	tr.Commit()
	println(tr.Proof())
	println(string(tr.Query([]byte{1, 2, 3})))
}
