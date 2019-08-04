package transaction

import (
	"sync"
)

type TxPool struct {
	M   map[string]struct{}
	Tx  []*Transaction
	Mtx sync.Mutex
}

// NewTxPool 初始化一个交易池
func NewTxPool() *TxPool {
	return &TxPool{
		M:  make(map[string]struct{}),
		Tx: make([]*Transaction, 0),
	}
}

// PutTx 存入一个交易
func (tp *TxPool) PutTx(tx *Transaction) {
	tp.Mtx.Lock()
	defer tp.Mtx.Unlock()
	tp.Tx = append(tp.Tx, tx)
	tp.M[tx.Thash] = struct{}{}
}

//FindTx 找放进区块的交易
func (tp *TxPool) FindTx() []*Transaction {
	tp.Mtx.Lock()
	defer tp.Mtx.Unlock()
	txs := tp.Tx[:4]
	tp.Tx = tp.Tx[4:]
	for _, v := range txs {
		delete(tp.M, v.Thash)
	}
	return txs
}
