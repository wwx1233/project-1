package block

type Block struct {
	Head BlockHead
	Body BlockBody
}

type BlockHead struct {
	PrevHash   []byte
	BlockHash  []byte
	MerkleTree []byte
	TimeStamp  int64
	Height     int
}
type BlockBody struct {
	NumberOfTrans int64
	Transactions  []*Transaction
}

func NewBlock() *Block {

	return
}
