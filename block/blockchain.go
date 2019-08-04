package block

type BlockChain struct {
	TopHash string
	Height  int
	//db      *bolt.DB
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		TopHash: "",
		Height:  0,
	}
}

// func NewBlockChain() BlockChain {

// 	return
// }
