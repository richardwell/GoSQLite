package sqlparser

//sql中建表drop类型
type drop struct {
	table string
}

//解析drop类型的sql语句
//返回drop类型指针
func parseDrop(sql string) (*drop, error) {
	return nil, nil
}

//返回类型
func (d *drop) GetType() int {
	return DROPTABLE
}

//返回表名
func (d *drop) GetTable() ([]string, error) {
	if d.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = d.table
	return str, nil
}

//不要实现
func (d *drop) GetAttribute() ([]string, error) {
	return nil, InvokeError
}

//不要实现
func (d *drop) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不要实现
func (d *drop) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不要实现
func (d *drop) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//不要实现
func (d *drop) GetConfig() ([][]string, error) {
	return nil, InvokeError
}
