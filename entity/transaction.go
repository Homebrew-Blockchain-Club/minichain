package entity

import (
	"github.com/Homebrew-Blockchain-Club/minichain/hasher"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Nonce    uint64
	GasLimit uint64
	From     []byte
	To       []byte
	Amount   uint64
	Data     []byte
	R, S, V  []byte
}

func Verify(tx Transaction) bool {
	r := tx.R
	s := tx.S
	v := tx.V
	tx.R = nil
	tx.S = nil
	tx.V = nil
	raw := typeconv.ToBytes(tx)
	hash := hasher.Hash(raw)
	sig := append(r, s...)
	sig = append(sig, v...)
	pubkey, _ := crypto.SigToPub(hash, sig)
	addr := crypto.PubkeyToAddress(*pubkey)
	for pos, x := range addr.Bytes() {
		if x != tx.From[pos] {
			return false
		}
	}
	return true
}
func Sign(tx *Transaction, privateKey []byte) {
	key, _ := crypto.ToECDSA(privateKey)
	txbyte := typeconv.ToBytes(*tx)
	hash := hasher.Hash(txbyte)
	sign, _ := crypto.Sign(hash, key)
	//r := new(big.Int).SetBytes(sign[:32])
	r := sign[:32]
	s := sign[32:64]
	v := sign[64:]
	tx.R = r
	tx.S = s
	tx.V = v
}
