package sqlparser

//sql中建表createindex类型
//create index indexName on T(id,name);
type createindex struct {
	name      string
	table     string
	attribude []string
}

//解析create index类型的sql语句
//参数：经过处理的正确sql语句
//返回：createtable类型指针
//注释：正确的sql如下
//create index indexName on T(id);
//create index indexName on T(id,name);
func parseCreateindex(sql string) (*createindex, error) {
	return nil, nil
}

//返回类型
func (c *createindex) GetType() int {
	return CREATINDEX
}

//返回表名 T
func (c *createindex) GetTable() ([]string, error) {
	if c.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = c.table
	return str, nil
}

//返回各列属性名 id/id，name,就是要索引的属性名
func (c *createindex) GetAttribute() ([]string, error) {
	return c.attribude, nil
}

//不实现
func (c *createindex) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (c *createindex) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不实现
func (c *createindex) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//create index类型，IndexName 表示索引名
func (c *createindex) GetConfig() (*Config, error) {
	f := &Config{}
	f.Arg3 = c.name
	return f, nil
}
