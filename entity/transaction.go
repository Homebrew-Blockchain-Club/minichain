package entity

import (
	"crypto/ecdsa"
	"math/big"

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
	R, S, V  *big.Int
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
	sig := append(r.Bytes(), s.Bytes()...)
	sig = append(sig, byte(v.Uint64()-27))
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
	key := typeconv.FromBytes[ecdsa.PrivateKey](privateKey)
	txbyte := typeconv.ToBytes(*tx)
	hash := hasher.Hash(txbyte)
	sign, _ := crypto.Sign(hash, &key)
	r := new(big.Int).SetBytes(sign[:32])
	s := new(big.Int).SetBytes(sign[32:64])
	v := big.NewInt(int64(sign[64] + 27))
	tx.R = r
	tx.S = s
	tx.V = v
}
