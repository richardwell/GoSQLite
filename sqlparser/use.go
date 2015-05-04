package sqlparser

type use struct {
	db string
}

//解析use类型的sql语句
//参数：经过处理的正确sql语句
//返回：use类型指针
//注释：正确的sql如下
//use dbname;
func parseUse(sql string) (*use, error) {
	return nil, nil
}

//返回类型
func (u *use) GetType() int {
	return USE
}

//不实现
func (u *use) GetTable() ([]string, error) {
	return nil, InvokeError
}

//不实现
func (u *use) GetAttribute() ([]string, error) {
	return nil, InvokeError
}

//不实现
func (u *use) GetAttributeType() ([]Type, error) {
	return nil, InvokeError
}

//不实现
func (u *use) GetValues() ([]interface{}, error) {
	return nil, InvokeError
}

//不实现
func (u *use) GetCondition() (*Where, error) {
	return nil, InvokeError
}

//不实现
func (u *use) GetConfig() (*Config, error) {
	f := &Config{}
	f.Arg3 = u.db
	return f, nil
}
