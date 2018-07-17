package logs

import (
	"fmt"
	"runtime"
	"time"
)

//type Logger interface {
//	Debug(string)
//	Error(string)
//	Info(string)
//	Log(string)
//	Warn(string)
//}

func Log(str ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("\n", now, file, line, str)
}
