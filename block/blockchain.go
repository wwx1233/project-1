package block

type BlockChain struct {
	TopHash []byte
	db      *bolt.DB
}

func NewBlockChain() BlockChain {

	return
}
