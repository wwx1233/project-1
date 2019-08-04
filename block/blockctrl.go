package block

import (
	"log"

	"github.com/LazyboyChen7/project/transaction"
)

// BlockCtrl 区块控制器
type BlockCtrl struct {
	Bc *BlockChain
	Tx *transaction.TxPool
}

// NewBlockCtrl 创建区块控制器
func NewBlockCtrl() *BlockCtrl {
	return &BlockCtrl{
		Bc: NewBlockChain(),
		Tx: transaction.NewTxPool(),
	}
}

// CheckBlock 检查是否满足打包条件
func (bc *BlockCtrl) CheckBlock() {

}

// CreateBlock 产生区块
func (bc *BlockCtrl) CreateBlock() *Block {
	prevHash := bc.Bc.TopHash
	height := bc.Bc.Height
	txs := bc.Tx.FindTx()
	block, err := NewBlock(prevHash, height, txs)
	if err != nil {
		log.Println(err)
		return nil
	}
	return block
}
