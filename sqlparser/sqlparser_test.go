package sqlparser

import (
	"testing"
)

//func Test_pretreatment(t *testing.T) {
//	sql := "create TABLE T(ID int primary key, name TEXT  )"
//	fmt.Println(pretreatment(sql))

//	sql = "  select *  from		T where id = 1   "
//	fmt.Println(pretreatment(sql))

//	sql = "  select T1 .id,T2. name  from T1,T2 where id = 1   "
//	fmt.Println(pretreatment(sql))

//	sql = "update T set name= ' tom ' where jidian =2. 3;"
//	fmt.Println(pretreatment(sql))

//	sql = "update T set name= ' tom ' where jidian in (2, 3. 0);"
//	fmt.Println(pretreatment(sql))

//	sql = "inst into T (ID,name ,jidian ) value (3,' TOM ',2.3  ) ;"
//	fmt.Println(pretreatment(sql))

//}

func Benchmark_pretreatment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sql := "inst into T (ID,name ,jidian ) value (3,' TOM ',2.3  ) ;"
		pretreatment(sql)
	}

}
