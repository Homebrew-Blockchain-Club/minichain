package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Homebrew-Blockchain-Club/minichain/comm"
	"github.com/Homebrew-Blockchain-Club/minichain/entity"
	"github.com/Homebrew-Blockchain-Club/minichain/typeconv"
	"github.com/stretchr/testify/assert"
)

func TestComm(t *testing.T) {
	// 创建Communicator实例
	communicator := comm.NewCommunicator()
	//test receive
	tx := entity.Transaction{}
	txData := typeconv.ToBytes(tx)
	pkg := comm.Package{
		Type: "transaction",
		Data: txData,
	}
	communicator.Receive(pkg)

	// 模拟账户地址
	address := []byte("mock-address")

	// 将地址转换为JSON格式
	addressJSON, err := json.Marshal(address)
	if err != nil {
		t.Fatalf("Failed to marshal address: %v", err)
	}

	// 创建一个HTTP请求
	req, err := http.NewRequest(http.MethodPost, "/query", bytes.NewBuffer(addressJSON))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 创建一个HTTP响应记录器
	w := httptest.NewRecorder()

	// 使用Communicator的gin引擎处理请求
	communicator.Router.ServeHTTP(w, req)

	// 检查响应状态码
	assert.Equal(t, http.StatusOK, w.Code)

	// 检查响应内容
	var acc entity.Account
	err = json.Unmarshal(w.Body.Bytes(), &acc)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 验证返回的账户信息
	expectedAccount := entity.Account{}
	assert.Equal(t, expectedAccount, acc)
}
