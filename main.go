package main

import (
	"log"

	"github.com/LazyboyChen7/project/agent"
)

func main() {
	log.Println("Process Start...")
	ag, err := agent.NewAgent()
	if err != nil {
		log.Println("create agent wrong...", err)
		return
	}
	log.Println("The wallet private key  is ", ag.Wlt.PrivateKey.D)
	log.Println("The wallet public key x is ", ag.Wlt.PrivateKey.PublicKey.X)
	log.Println("The wallet public key y is ", ag.Wlt.PrivateKey.PublicKey.Y)
	log.Println("The server port is ", ag.Svr.Port)
	log.Println("The blockchain height is ", ag.Blc.Bc.Height)
	log.Println("The blockchain top hash is ", ag.Blc.Bc.TopHash)

}
