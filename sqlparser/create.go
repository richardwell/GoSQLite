package sqlparser

//sql中建表create类型
type create struct {
	table      string
	attribude  []string
	attriType  []Type
	primaryKey []string
	indext     []string
}

//解析create类型的sql语句
//返回ctreate类型指针
func parseCreate(sql string) (*create, error) {
	return nil, nil
}

//返回类型
func (c *create) GetType() int {
	return CREATETABLE
}

//返回表名
func (c *create) GetTable() ([]string, error) {
	if c.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = c.table
	return str, nil
}

//返回各列属性名
func (c *create) GetAttribute() ([]string, error) {
	return c.attribude, nil
}

//返回各列属性对应类型
func (c *create) GetAttributeType() ([]Type, error) {
	return c.attriType, nil
}

//不要实现
func (c *create) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不要实现
func (c *create) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//create类型，PrimaryKey 表示主码属性列，Index 表示要索引的列
func (c *create) GetConfig() (*Config, error) {
	f := &Config{}
	f.PrimaryKey = c.primaryKey
	f.Index = c.indext
	return f, nil
}
