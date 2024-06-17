// 本包应当对这些函数用gin进行封装以提供http调用
// 本包应使用/exec的结构体和函数
package comm

import "sync"

type Communicator struct {
}

// 创建新的交流器
// 创建一个gin的goroutine以接收http请求、绑定本包的函数调用
func NewCommunicator() Communicator {
	return Communicator{}
}

var SendMutex sync.Mutex

// 发送包，这个包可以包括一笔交易，也可以是一个新区块
// 请注意这个函数要使用mutex锁
func (*Communicator) Send(Package) {

}

var ReceiveMutex sync.Mutex

// 接收包
// 此函数要使用mutex锁
func (*Communicator) Receive(Package) {
}

// 假如本包是一个交易包，应当进行的操作
func TransactionReceived(Package) {}

// 假如本包是一个区块包，应当进行的操作
func BlockReceived(Package) {

}
