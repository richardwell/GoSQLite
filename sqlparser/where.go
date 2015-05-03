package sqlparser

//sql中where那部分语句，仅支持：
//运算符：=,!=,<,<=,>,>=
//IN 嵌套
//ALL 嵌套
//NOT
//AND
//OR
type Where struct {
	//属性名
	Selects []string

	//运算符
	Operation []func() bool

	//条件的值 如 =？ 中的？
	SelectArgs []interface{}
}

//只传入sql语句中where那部分，如 where 1！=2
func parseWhere(sql string) (*Where, error) {
	return nil, nil
}
