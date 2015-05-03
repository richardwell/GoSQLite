package sqlparser

//sql中建表update类型
type update struct {
	table     string
	attribude []string
	values    []interface{}
	where     *Where
}

//解析update类型的sql语句
//返回update类型指针
func parseUpdate(sql string) (*update, error) {
	return nil, nil
}

//返回类型
func (u *update) GetType() int {
	return UPDATE
}

//返回表名
func (u *update) GetTable() ([]string, error) {
	if u.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = u.table
	return str, nil
}

//返回各列属性名
func (u *update) GetAttribute() ([]string, error) {
	return u.attribude, nil
}

//不要实现
func (u *update) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//返回各列属性对应的值
func (u *update) GetValues() ([]interface{}, error) {
	return u.values, nil
}

//where中的条件
func (u *update) GetCondition() (*Where, error) {
	return u.where, nil
}

//不要实现
func (u *update) GetConfig() ([][]string, error) {
	return nil, InvokeError
}
