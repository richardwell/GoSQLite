package sqlparser

import (
	"errors"
)

//sql语句类型标识
const (
	CREATETABLE = iota
	CREATINDEX
	DROPTABLE
	DROPINDEX
	SELECT
	INSERT
	UPDATE
	DELETE
)

//数据类型标识
const (
	//INT(size) -2147483648 到 2147483647 常规。在括号中规定最大位数。
	INT = 20 + iota

	//BIGINT(size) -9223372036854775808 到 9223372036854775807 常规。在括号中规定最大位数。
	BIGINT

	//FLOAT(size,d)	带有浮动小数点的小数字。在括号中规定最大位数。在 d 参数中规定小数点右侧的最大位数。
	FLOAT

	//TEXT 存放最大长度为 65,535 个字符的字符串。
	TEXT

	//VARCHAR(size)	保存可变长度的字符串（可包含字母、数字以及特殊字符）。
	//在括号中指定字符串的最大长度。最多 255 个字符。
	//注释：如果值的长度大于 255，则被转换为 TEXT 类型。
	VARCHAR

	//BLOB	用于 BLOBs (Binary Large OBjects)。存放最多 65,535 字节的数据。
	BLOB

	//DATE() 日期。格式：YYYY-MM-DD 注释：支持的范围是从 '1000-01-01' 到 '9999-12-31'
	DATE

	//TIME()	时间。格式：HH:MM:SS 注释：支持的范围是从 '-838:59:59' 到 '838:59:59'
	TIME

	//DATETIME()	*日期和时间的组合。格式：YYYY-MM-DD HH:MM:SS
	//注释：支持的范围是从 '1000-01-01 00:00:00' 到 '9999-12-31 23:59:59'
	DATETIME
)

//sql关键字,可选
const (
	CREATE_TABLE = "create table"
)

//定义一些通用的错误类型
var (
	InvokeError   error = errors.New("错误，你不能调用这个方法")
	NilTableError error = errors.New("错误，表名为空")
)
