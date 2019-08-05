package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LazyboyChen7/project/agent"
	"github.com/LazyboyChen7/project/transaction"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("\n%s \n\n", "Usage: go run main.go [port].")
		return
	}
	log.Println("Process Start...")
	ag, err := agent.NewAgent(os.Args[1])
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

	go ag.Start()

	tx := transaction.CreateTx("a", "a", "a")
	fmt.Println(tx.Thash)
}
