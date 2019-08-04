package transaction

import "sync"

type TxPool struct {
	M   map[string]struct{}
	Tx  []*Transaction
	Mtx sync.Mutex
}

// CheckBlock 检查是否满足打包条件
func CheckBlock() {

}

// FindTx 找放进区块的交易
func FindTx() []Transaction {
	Mtx.Lock()
	defer Mtx.Unlock()
	txs := Tx[:4]
	Tx = Tx[4:]
	for _,v := txs {
		delete(M, v)
	}
	return txs
}

func PutTx(tx *Transaction) {
	agent.TxPool.Mtx.Lock()
	defer agent.TxPool.Mtx.Unlock()
	Tx = append(Tx, tx)
	hash := hash(tx)
	M[] = struct{}
}