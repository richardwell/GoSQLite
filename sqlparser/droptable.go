package sqlparser

//sql中dropt able类型
type droptable struct {
	table string
}

//解析dropt able类型的sql语句
//参数：经过处理的正确sql语句
//返回：droptable类型指针
//注释：正确的sql如下
//drop table T;
func parseDrop(sql string) (*droptable, error) {
	return nil, nil
}

//返回类型
func (d *droptable) GetType() int {
	return DROPTABLE
}

//返回表名 T
func (d *droptable) GetTable() ([]string, error) {
	if d.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = d.table
	return str, nil
}

//不要实现
func (d *droptable) GetAttribute() ([]string, error) {
	return nil, InvokeError
}

//不要实现
func (d *droptable) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (d *droptable) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不实现
func (d *droptable) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//不实现
func (d *droptable) GetConfig() ([][]string, error) {
	return nil, InvokeError
}
