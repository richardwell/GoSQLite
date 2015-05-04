package sqlparser

//sql中建表insert类型
type insert struct {
	table     string
	attribude []string
	values    []interface{}
}

//解析insert类型的sql语句
//参数：经过处理的正确sql语句
//返回:insert类型指针
//注释：正确sql语句
//insert into T(id,name,info) value(1,'tom','babab');
//insert into T(id) value(2);
func parseInsert(sql string) (*insert, error) {
	return nil, nil
}

//返回类型
func (i *insert) GetType() int {
	return INSERT
}

//返回表名
func (i *insert) GetTable() ([]string, error) {
	if i.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = i.table
	return str, nil
}

//返回各列属性名
func (i *insert) GetAttribute() ([]string, error) {
	return i.attribude, nil
}

//不实现
func (i *insert) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//返回各列属性对应的值
func (i *insert) GetValues() ([]interface{}, error) {
	return i.values, nil
}

//不实现
func (i *insert) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//不实现
func (i *insert) GetConfig() (*Config, error) {
	return nil, InvokeError
}
