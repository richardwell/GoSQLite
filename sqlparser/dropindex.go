package sqlparser

//sql中dropt index类型
type dropindex struct {
	name  string
	table string
}

//解析dropt able类型的sql语句
//参数：经过处理的正确sql语句
//返回：droptable类型指针
//注释：正确的sql如下
//drop index indexName on T;
func parseDropindex(sql string) (*dropindex, error) {
	return nil, nil
}

//返回类型
func (d *dropindex) GetType() int {
	return DROPTABLE
}

//返回表名 T
func (d *dropindex) GetTable() ([]string, error) {
	if d.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = d.table
	return str, nil
}

//不要实现
func (d *dropindex) GetAttribute() ([]string, error) {
	return nil, InvokeError
}

//不要实现
func (d *dropindex) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (d *dropindex) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不实现
func (d *dropindex) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//drop index类型，IndexName 表示索引名 indexName
func (d *dropindex) GetConfig() (*Config, error) {
	f := &Config{}
	f.Arg3 = d.name
	return f, nil
}
