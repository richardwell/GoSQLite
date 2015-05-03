//该包的主流程控制
package sqlparser

//下层调用的函数
//参数是从客户端得到
//返回解析结果接口
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
//参数从客户端得到
//返回经过处理的
func pretreatment(sql string) string {
	return ""
}

//sql语句检错
//参数应从 pretreatment 函数得到
//返回正确的sql语句或抛出错误
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
