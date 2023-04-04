package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

var (
	Logger *logging.Logger
	format logging.Formatter
)

func InitLogger() {
	//time.Now().In(time.FixedZone("UTC", 8*3600))
	tid := uuid.NewV4().String()
	if Logger == nil {
		Logger = logging.MustGetLogger("example")
	}
	format = logging.MustStringFormatter(`{"id":"%{id:03x}","function":"%{shortfile}","class":"%{shortfunc}","type":"%{level:.4s}","timestamp":"%{time:2006/01/02 15:04:05.000}","data":"%{message}","tid":"` + fmt.Sprint(tid) + `"}`)
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1 := logging.NewLogBackend(FileLogger(), "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatter, backend2Formatter)
}

func FileLogger() *os.File {

	_, err := os.Stat("./runtime") //os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			os.MkdirAll("runtime", os.ModePerm)
		}
	}
	file, error := os.OpenFile("runtime/go_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if error != nil {
		fmt.Println(error.Error())
		os.Exit(1)
	}
	return file
}

func Blogs() {
	f := &logs.PatternLogFormatter{
		Pattern:    "{'file':'%F:%n','time':'%w','type':'%t','info':'%m'}",
		WhenFormat: "2006-01-02 15:04:05",
	}
	logs.RegisterFormatter("pattern", f)
	_ = logs.SetGlobalFormatter("pattern")
	logs.Info("hello, world")

}
