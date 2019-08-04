package block

import (
	"time"

	"github.com/LazyboyChen7/project/transaction"
)

// Block 区块
type Block struct {
	Head *BlockHead
	Body *BlockBody
}

// BlockHead 区块头
type BlockHead struct {
	PrevHash   string `json:"prevhash"`
	BlockHash  string `json:"blockhash"`
	MerkleTree []byte `json:"merkletree"`
	TimeStamp  int64  `json:"timestamp"`
	Height     int    `json:"height"`
}

// BlockBody 区块体
type BlockBody struct {
	NumberOfTrans int64
	Transactions  []*transaction.Transaction
}

// NewBlock 新打包一个区块
func NewBlock(prevHash []byte, height int) (*Block, error) {
	txs := FindTx()
	merkleroot, err := GetMerkleRoot(Txs)
	if err != nil {
		return nil, err
	}
	createTime := time.Now().UnixNano()
	blockhash := utils.GetBlockHash(prevHash, merkleroot, createTime, height)
	blkHead := &BlockHead{
		PrevHash:   prevHash,
		BlockHash:  blockhash,
		MerkleTree: merkleroot,
		TimeStamp:  createTime,
		Height:     height + 1,
	}
	blkBody := &BlockBody{
		NumberOfTrans: len(txs),
		Transactions: txs
	}
	return &Block{
		Head: blkHead,
		Body: blkBody,
	}
}
