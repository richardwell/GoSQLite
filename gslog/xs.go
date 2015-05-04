package gslog

import (
	"log"
	"os"
	"path"
)

var (
	xslog    = newxslog()                                 //xs的日志
	debug    = true                                       //日志输出开关
	fileName = "C:/Users/Administrator/Desktop/xslog.log" //日志路径
	prex     = "xs: "                                     //前缀
)

func newxslog() *log.Logger {
	f, err := os.OpenFile(fileName, os.O_APPEND, os.ModeSetuid)
	if err != nil {
		os.MkdirAll(path.Dir(fileName), os.ModeDir)
		os.Create(fileName)
		t, _ := os.OpenFile(fileName, os.O_APPEND, os.ModeSetuid)
		f = t
	}
	//defer f.Close()
	l := log.New(f, prex, log.Ldate|log.Ltime|log.Lshortfile)
	return l
}

func XSPrintln(v ...interface{}) {
	if debug {
		xslog.Println(v, "\r\n")
	}
}
