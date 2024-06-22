package testing

import (
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestTransactionSignature(t *testing.T) {

	privatekey, _ := crypto.GenerateKey()
	publickey := privatekey.PublicKey
	address := crypto.PubkeyToAddress(publickey)
	a := entity.Transaction{
		From: address[:],
	}
	privkeybyte := crypto.FromECDSA(privatekey)
	entity.Sign(&a, privkeybyte)
	if !entity.Verify(a) {
		panic("entity doesn't match")
	}
}
