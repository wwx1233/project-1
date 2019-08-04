package block

import (
	"bytes"
	"time"

	"github.com/LazyboyChen7/project/transaction"
	"github.com/LazyboyChen7/project/utils"
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
	MerkleTree string `json:"merkletree"`
	TimeStamp  int64  `json:"timestamp"`
	Height     int    `json:"height"`
}

// BlockBody 区块体
type BlockBody struct {
	NumberOfTrans int
	Transactions  []*transaction.Transaction
}

// NewBlock 新打包一个区块
func NewBlock(prevHash string, height int, txs []*transaction.Transaction) (*Block, error) {
	merkleroot, err := GetMerkleRoot(txs)
	if err != nil {
		return nil, err
	}
	createTime := time.Now().UnixNano()
	blockhash := GetBlockHash(prevHash, merkleroot, createTime, height)
	blkHead := &BlockHead{
		PrevHash:   prevHash,
		BlockHash:  blockhash,
		MerkleTree: merkleroot,
		TimeStamp:  createTime,
		Height:     height + 1,
	}
	blkBody := &BlockBody{
		NumberOfTrans: len(txs),
		Transactions:  txs,
	}
	return &Block{
		Head: blkHead,
		Body: blkBody,
	}, nil
}

// GetBlockHash 获取区块hash
func GetBlockHash(prehash, root string, tmp int64, height int) string {
	var buf bytes.Buffer
	buf.WriteString(prehash)
	buf.WriteString(root)
	buf.Write(utils.Int64ToByte(tmp))
	buf.Write(utils.IntToByte(height))
	return utils.Hash(buf.Bytes())
}

// GetMerkleRoot 获取Merkle根
func GetMerkleRoot(txs []*transaction.Transaction) (string, error) {

	return "", nil
}
