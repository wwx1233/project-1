package transaction

import "project/db"

type Transaction struct {
	IsData         bool    // 创建物品/转移物品
	Data           db.Meta // 物品元信息
	CreateInstance []byte  // 物品初始归属
	Vin            []byte  // 流转输出方
	Vout           []byte  // 流传输入方
}
