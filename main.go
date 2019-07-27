package main

import (
	"fmt"
	"project/wallet"
)

var MyWallet *wallet.Wallet

func main() {
	MyWallet := wallet.NewWallet()
	//server := Newserver()
	fmt.Println(MyWallet)

	// cli.Cli{}
	// cli.Run()
}
