package sqlparser

//sql中建表select类型
type select_ struct {
	table      []string
	attribute  []string
	where      *Where
	orderAttri string
	isDesc     bool
	isDistinct bool
}

//解析select类型的sql语句
//参数：经过处理的正确sql语句
//返回:select类型指针
//注释1:正确sql
//select * from T;
//select * from T where id=1001 and name='TOM';
//select T1.id T2.name from T1,T2 where T1.id NOT IN (select id from T2) order by T1.id desc;
//注释2:where后面的语句不在这里解析，即 where ... order by 之间的内容给 parseWhere() 解析，
//传给parseWhere 的语句要包含 where 即：where T1.id NOT IN (select id from T2)
func parseSelect(sql string) (*select_, error) {
	return nil, nil
}

//返回类型
func (s *select_) GetType() int {
	return SELECT
}

//返回表名 T/ T1,T2
func (s *select_) GetTable() ([]string, error) {
	if s.table == nil {
		return nil, NilTableError
	}
	return s.table, nil
}

//返回各列属性名
func (s *select_) GetAttribute() ([]string, error) {
	return s.attribute, nil
}

//不实现
func (s *select_) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (s *select_) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//where中的条件
func (s *select_) GetCondition() (*Where, error) {
	return s.where, nil
}

//select类型，IsDistinct 是否去掉重复元组，IsDesc 表示排序方式,OrderAttri根据哪个属性排序
func (s *select_) GetConfig() (*Config, error) {
	f := &Config{}
	f.Arg0 = s.isDistinct
	f.Arg1 = s.isDesc
	f.Arg3 = s.orderAttri
	return f, nil
}
