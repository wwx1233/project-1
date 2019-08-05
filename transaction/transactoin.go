package transaction

import (
	"bytes"

	"github.com/LazyboyChen7/project/utils"
)

type Transaction struct {
	// CreateInstance []byte // 物品初始归属
	// Case []int // 案件信息
	Data  string // 物品元信息
	Vin   string // 流转输出方
	Vout  string // 流传输入方
	Thash string
}

// CreateTx 创建交易
func CreateTx(data string, from, to string) Transaction {
	return Transaction{
		Data:  data,
		Vin:   from,
		Vout:  to,
		Thash: Txhash(data, from, to),
	}
}

// Verify 验证交易
func VerifyTx(tx *Transaction) bool {
	// 1，验证交易
	if tx.Vin == "" {
		//data => vout
		return true
	} else {
		// 查数据库
		// if vin have data {
		// 	vin delete data
		// 	vout put data
		// 	return true
		// } else {
		// 	return false
		// }
	}

	return true
}

func Txhash(d, f, t string) string {
	var bf bytes.Buffer
	bf.WriteString(d)
	bf.WriteString(f)
	bf.WriteString(t)
	bt := bf.Bytes()
	return utils.Hash(bt)
}
