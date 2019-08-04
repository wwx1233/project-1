package db

type Meta struct {
	Name        []byte // 物品名称
	CraetedTime int64  // 入库时间
	Region      []byte // 存储位置
	LawCase     []byte // 所属案件
}

type CreateMeta interface{}

func NewMetaDate(name, region, lawcase []byte, createdtime int64) *Meta {

	return nil
}
