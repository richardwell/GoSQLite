//该包的主流程控制
package sqlparser

//下层调用的函数
//参数:是从客户端得到一条字符串
//返回值：解析得到的结果接口
func ParseSQL(sql string) (*Parser, error) {
	pres := pretreatment(sql)
	dets, err := detect(pres)
	if err != nil {
		//???
	}
	par, err := parse(dets)
	if err != nil {
		//????
	}
	return par, nil
}

//sql语句预处理,删除多余空格，全小写，加；
//参数：从客户端得到
//返回：经过处理的sql字符串
//注释：将sql中的关键字，数据类型，表名，属性名等全改成小写
//删除多余空格，如 create table T ( id int),要删成 create table T(id int),
//弄成单词与单词之间一个空格，单词与符号之间没有空格
//如果最后一个不是';',则加上
func pretreatment(sql string) string {
	return ""
}

//sql语句检错
//参数：应从 pretreatment 函数得到
//返回：正确的sql语句或抛出错误
func detect(sql string) (string, error) {
	return "", nil
}

//解析sql语句，主要分发给sql类型解析器
//参数应从 detect 函数得到
//返回解析结果接口
func parse(sql string) (*Parser, error) {
	//掉用 parseXXX（）函数
	return nil, nil
}
