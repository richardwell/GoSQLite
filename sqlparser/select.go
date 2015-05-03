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
//返回select类型指针
func parseSelect(sql string) (*select_, error) {
	return nil, nil
}

//返回类型
func (s *select_) GetType() int {
	return SELECT
}

//返回表名
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

//不要实现
func (s *select_) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不要实现
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
	f.IsDesc = s.isDesc
	f.IsDistinct = s.isDistinct
	f.OrderAttri = s.orderAttri
	return f, nil
}
