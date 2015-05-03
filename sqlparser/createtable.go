package sqlparser

//sql中建表createtable类型
//create table T(id int(10));
type createtable struct {
	table      string
	attribude  []string
	attriType  []Type
	primaryKey []string
}

//解析create table类型的sql语句
//参数：经过处理的正确sql语句
//返回：ctreate类型指针
//注释：正确的sql如下
//create table T(id int(10) primary key,name varchar(50),info text);	//primary key 可以省略
//create table T(id int(10) primary key,name varchar(50) primary key,info text);	//primary key 可以多个
//create类型中 attriType 应这样实现
//type_id:=Type{Type:func()int{return INT},Size:func()int{return 10}}
//就是Type里面的函数在这里写
func parseCreate(sql string) (*createtable, error) {
	return nil, nil
}

//返回类型
func (c *createtable) GetType() int {
	return CREATETABLE
}

//返回表名 T
func (c *createtable) GetTable() ([]string, error) {
	if c.table == "" {
		return nil, NilTableError
	}
	str := make([]string, 1, 1)
	str[0] = c.table
	return str, nil
}

//返回各列属性名 id，name，info
func (c *createtable) GetAttribute() ([]string, error) {
	return c.attribude, nil
}

//返回各列属性对应类型
func (c *createtable) GetAttributeType() ([]Type, error) {
	return c.attriType, nil
}

//不实现
func (c *createtable) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不实现
func (c *createtable) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//createtable类型，PrimaryKey 表示主码属性列
func (c *createtable) GetConfig() (*Config, error) {
	f := &Config{}
	f.PrimaryKey = c.primaryKey
	return f, nil
}
