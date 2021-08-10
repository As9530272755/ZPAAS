package log

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// 通过闭包的手段
// 编写日志包装函数，等会在路由中调用该函数，因为我们的 GetUser AddUser DeleteUser Edit 等多个函数满足 http.HandlerFunc 方法
func LoggerWrapper(action http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		action(rw, r)

		// 判断日志文件是否存在
		logFile := "weblog.log"
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 写入日志文件
		log.SetOutput(file)
		log.SetFlags(log.Flags() | log.Lshortfile)
		log.Printf(`[%s] %s %s %s`, r.RemoteAddr, r.Method, r.URL.String(), r.Header.Get("User-Agent"))
	}
}
