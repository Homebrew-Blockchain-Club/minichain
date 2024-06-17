package minichain

import (
	"github.com/Homebrew-Blockchain-Club/minichain/cli"
	"github.com/Homebrew-Blockchain-Club/minichain/comm"
)

func main() {
	cli.NewCLI()
	comm.NewCommunicator()
}
