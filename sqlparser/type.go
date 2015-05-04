package sqlparser

//select，update 等结构体要实现的接口
//向下提供的解析结果接口
type Parser interface {
	//返回类型
	GetType() int

	//返回表名
	//除了select要多表查询时返回多表外
	//其他sql类型返回的表名放在s[0]中
	GetTable() ([]string, error)

	//返回各列属性名
	GetAttribute() ([]string, error)

	//返回各列属性对应类型
	GetAttributeType() ([]Type, error)

	//返回各列属性对应的值
	GetValues() ([]interface{}, error)

	//where中的条件
	GetCondition() (*Where, error)

	//某些类型的额外值
	GetConfig() (*Config, error)
}

//数据类型，包含所有类型
type Type struct {
	//返回类型
	Type func() int

	//返回大小，如INT(size) 中的size，-1表示不限定
	Size func() int
}

//如果是use类型，Arg3 表示数据库名
//如果是create table类型，Arg2 表示主码属性列
//如果是create index类型，Arg3 表示索引名
//如果是drop index类型，Arg3 表示索引名
//如果是select类型，Arg0 是否去掉重复元组，Arg1 表示排序方式, Arg3 表示根据哪个属性排序
type Config struct {
	Arg0 bool
	Arg1 bool
	Arg2 []string
	Arg3 string
}
