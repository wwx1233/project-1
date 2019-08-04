package transaction

import "github.com/LazyboyChen7/project/agent"

type Transaction struct {
	Data string // 物品元信息
	// CreateInstance []byte // 物品初始归属
	// Case []int // 案件信息
	Vin  []byte // 流转输出方
	Vout []byte // 流传输入方
}

// CreateTx 创建交易
func (a *agent.Agent) CreateTx(data string, addr []byte) *Transaction {
	return &Transaction{
		Data: data,
		Vin:  a.Wlt.Address,
		Vout: addr,
	}
}

// Verify 验证交易
func VerifyTx(tx *Transaction) bool {
	// 1，验证交易
	vin := tx.Vin
	vout := tx.Vout
	data := tx.Data
	if vin == "" {
		//data => vout
		return true
	} else {
		// 查数据库
		if vin have data {
			vin delete data
			vout put data
			return true
		} else {
			return false
		}
	}
}
