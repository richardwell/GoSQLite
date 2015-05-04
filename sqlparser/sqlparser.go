//该包的主流程控制
package sqlparser

import (
	"GoSQLite/gslog"
	"regexp"
	"strings"
)

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
	sql = strings.ToLower(sql)
	re, err := regexp.Compile("\\s+")
	if err != nil {
		gslog.XSPrintln("ERROR", err)
	}
	sql = re.ReplaceAllLiteralString(sql, " ")
	sql = strings.TrimSpace(sql)
	reop, err := regexp.Compile("[()\\-\\.,='*;]")
	if err != nil {
		gslog.XSPrintln("ERROR", err)
	}
	rezm, err := regexp.Compile("[1-9a-z]")
	if err != nil {
		gslog.XSPrintln("ERROR", err)
	}
	len_ := len(sql)
	isTrim_f = func(i int) bool {
		l := le
		j := i - 1
		if j < 0 {
			return true
		}
		k := i + 1
		if k >= len_ {
			return true
		}
		b1 := reop.MatchString(sql[j:j+1]) && reop.MatchString(sql[k:k+1])
		b2 := rezm.MatchString(sql[j:j+1]) && reop.MatchString(sql[k:k+1])
		b3 := reop.MatchString(sql[j:j+1]) && rezm.MatchString(sql[k:k+1])
		if b1 || b2 || b3 {
			return true
		}
		return false
	}
	bsql := []byte{}
	for i := 0; i < len_; i++ {
		if sql[i] == " " {
			if isTrim_f(i) {
				continue
			}
		}
		bsql = append(bsql, s[i])
	}
	sql = string(bsql)
	return sql
}

//sql语句检错
//参数：应从 pretreatment 函数得到
//返回：正确的sql语句或抛出错误
func detect(sql string) (string, error) {
	//reop, err := regexp.Compile("[^1-9a-z()\\-\\.,='*; ]")
	return "", nil
}

//解析sql语句，主要分发给sql类型解析器
//参数应从 detect 函数得到
//返回解析结果接口
func parse(sql string) (*Parser, error) {
	//掉用 parseXXX（）函数
	return nil, nil
}
