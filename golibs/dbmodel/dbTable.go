package dbmodel

type DbTableInterface interface {
	//表名
	TableName() string
	//表引擎
	TableEngine() string
	//边编码
	TableEncode() string
	//表注释
	TableComment() string
	//单字段索引
	TableIndex() [][]string
	//多字段索引
	TableUnique() [][]string
}

type DbTable struct {
}

func (this *DbTable) TableEngine() string {
	return "Innodb"
}

func (this *DbTable) TableEncode() string {
	return "utf8"
}
func (this *DbTable) TableComment() string {
	return ""
}
