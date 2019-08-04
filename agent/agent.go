package agent

import (
	"log"

	"github.com/LazyboyChen7/project/block"
	"github.com/LazyboyChen7/project/server"
	"github.com/LazyboyChen7/project/wallet"
)

// Agent 执行实体
type Agent struct {
	Wlt *wallet.Wallet
	Svr *server.Server
	Blc *block.BlockCtrl
}

// NewAgent 初始化Agent
func NewAgent() (*Agent, error) {
	log.Println("Create new Agent now")
	wlt, err := wallet.NewWallet()
	if err != nil {
		log.Println("-- Craete new wallet wrong:  ", err)
		return nil, err
	}
	log.Println("-- Craete new wallet Success!")
	svr, err := server.NewServer()
	if err != nil {
		log.Println("-- Craete new server wrong:  ", err)
		return nil, err
	}
	log.Println("-- Craete new server Success!")
	blc := block.NewBlockCtrl()
	if err != nil {
		log.Println("-- Craete new blockcontroller wrong:  ", err)
		return nil, err
	}
	log.Println("-- Craete new blockcontroller Success!")
	log.Println("Create new Agent success !!!")
	return &Agent{
		Wlt: wlt,
		Svr: svr,
		Blc: blc,
	}, nil
}
