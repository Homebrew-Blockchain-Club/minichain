package comm

type Communicator struct {
}

func NewCommunicator() Communicator {
	return Communicator{}
}
func (*Communicator) Send(Package) {

}

func (*Communicator) Receive(Package) {
}

func TransactionReceived(Package) {}
func BlockReceived(Package) {

}
