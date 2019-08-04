package agent

import (
	"github.com/LazyboyChen7/project/server"
	"github.com/LazyboyChen7/project/wallet"
)

// Agent 执行实体
type Agent struct {
	Wlt    wallet.Wallet
	Svr    server.Server
	TxPool transaction.TxPool //交易池
}
