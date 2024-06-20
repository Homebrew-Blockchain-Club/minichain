package ds

type AbstractMPT interface {
	Commit()
	Store(key, val []byte)
	Query([]byte) []byte
	Proof() bool
}
