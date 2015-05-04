package sqlparser

//sql中建表delete类型
type delete_ struct {
	table string
	where *Where
}

//解析delete类型的sql语句
//参数：经过处理的正确sql语句
//返回：delete类型指针
//注释：正确sql语句
//delete from T where id=3;
//where 语句扔给 parseWhere() 解析，这里是 where id=3
func parseDelete(sql string) (*delete_, error) {
	return nil, nil
}

//返回类型
func (d *delete_) GetType() int {
	return DELETE
}

//返回表名
func (d *delete_) GetTable() ([]string, error) {
	if d.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = d.table
	return str, nil
}

//不实现
func (d *delete_) GetAttribute() ([]string, error) {
	return nil, InvokeError
}

//不实现
func (d *delete_) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (d *delete_) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//where中的条件
func (d *delete_) GetCondition() (*Where, error) {
	return d.where, nil
}

//不实现
func (d *delete_) GetConfig() (*Config, error) {
	return nil, InvokeError
}
