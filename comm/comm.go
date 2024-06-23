// 本包应当对这些函数用gin进行封装以提供http调用
// 本包应使用/exec的结构体和函数
// 在项目的第一阶段 Send函数收到新的交易只会调用本地/exec下的函数 收到新的区块，也只会直接调用本地的/exec下的函数进行验证和加入
package comm

import (
	"container/list"
	"log"
	"sync"

	"github.com/Homebrew-Blockchain-Club/minichain/ds"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/exec"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/gin-gonic/gin"
)

type Communicator struct {
	Router *gin.Engine
	ctrl   exec.AbstractController
}

const CONTROLLER_UNIMPLEMENTED = true

type MockController struct {
	Transactions []entity.Transaction
	Blocks       []ds.Block
}

func (m *MockController) AddTransaction(tx entity.Transaction) {
	m.Transactions = append(m.Transactions, tx)
	log.Println("Mock AddTransaction called with:", tx)
}

func (m *MockController) AddBlock(block ds.Block) {
	m.Blocks = append(m.Blocks, block)
	log.Println("Mock AddBlock called with:", block)
}

func (m *MockController) QueryAccount([]byte) entity.Account {
	return entity.Account{}
}

// 创建新的交流器
// 创建一个gin的goroutine以接收http请求、绑定本包的函数调用
func NewCommunicator() Communicator {
	Router := gin.Default()
	var ctrl exec.AbstractController
	if !CONTROLLER_UNIMPLEMENTED {
		ctrl = exec.NewController()
	} else {
		ctrl = &MockController{}
	}
	comm := Communicator{
		Router: Router,
		ctrl:   ctrl,
	}
	// 设置路由
	Router.POST("/receive", func(c *gin.Context) {
		var pkg Package
		if err := c.ShouldBindJSON(&pkg); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		comm.Receive(pkg)
		c.JSON(200, gin.H{"status": "received"})
	})

	Router.POST("/query", func(c *gin.Context) {
		var address []byte
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		acc := comm.Query(address)
		c.JSON(200, acc)
	})
	// 启动 HTTP 服务
	go func() {
		if err := Router.Run(); err != nil {
			panic(err)
		}
	}()

	return comm

}

var SendMutex sync.Mutex
var Sendqueue list.List

// 发送包，这个包可以包括一笔交易，也可以是一个新区块
// 请注意这个函数要使用mutex锁
func (*Communicator) Send(Package) {

}

var ReceiveMutex sync.Mutex
var Receivequeue list.List

func (comm *Communicator) Query(address []byte) entity.Account {
	return comm.ctrl.QueryAccount(address)
}

// 接收包
// 此函数要使用mutex锁
func (comm *Communicator) Receive(p Package) {
	// 将请求放入队列
	Receivequeue.PushBack(&p)
	ReceiveMutex.Lock()
	defer ReceiveMutex.Unlock()
	// 处理队列中的请求

	element := Receivequeue.Front()
	pkg := element.Value.(*Package)
	Receivequeue.Remove(element)

	// 根据包的类型处理请求
	switch pkg.Type {
	case "transaction":
		var transaction entity.Transaction
		transaction = typeconv.FromBytes[entity.Transaction](pkg.Data)
		comm.ctrl.AddTransaction(transaction)
	case "block":
		var block ds.Block
		block = typeconv.FromBytes[ds.Block](pkg.Data)
		comm.ctrl.AddBlock(block)
	}

}
