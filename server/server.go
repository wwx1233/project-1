package server

import (
	"log"
	"project/wallet"
)

type IdenMsg struct {
	Addr   []byte
	PubKey []byte
}

type Server struct {
	address string
	Wallet  wallet.Wallet
	IdenSet map[string]IdenMsg
}

const (
	SuperNode = ":9999"
)

func NewServer() *Server {
	MyWallet, MyAddr, err := wallet.NewWallet()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		address: MyAddr,
		Wallet:  MyWallet,
		IdenSet: make(map[string]IdenMsg),
	}
}
