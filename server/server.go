package server

import (
	"github.com/LazyboyChen7/project/utils"
)

type IdenMsg struct {
	Addr   []byte
	PubKey []byte
}

type Server struct {
	port    string
	IdenSet map[string]IdenMsg
}

const (
	SuperNode = ":9999"
)

func NewServer() *Server {
	port := utils.GetValidPort()
	return &Server{
		port:    port,
		IdenSet: make(map[string]IdenMsg),
	}
}
