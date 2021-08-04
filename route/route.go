package route

import (
	"exec/controllers"
	"exec/log"
	"net/http"
)

func Register() {
	// GetUser AddUser DeleteUser Edit 这几个函数满足 http.HandlerFunc 方法所以能够实现 log.LoggerWrapper() 的调用关系
	http.HandleFunc("/", log.LoggerWrapper(controllers.GetUser))
	http.HandleFunc("/create/", log.LoggerWrapper(controllers.AddUser))
	http.HandleFunc("/delete/", log.LoggerWrapper(controllers.DeleteUser))
	http.HandleFunc("/edit/", log.LoggerWrapper(controllers.Edit))
}
