package ds

type MPT struct {
}

func NewRoot() *MPT {
	return nil
}

func (tr *MPT) Store([][]byte) {

}
func (tr *MPT) Query([]byte) {

}
func (tr MPT) Proof() bool {
	return true
}
