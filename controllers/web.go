package controllers

import (
	"exec/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// 记录客户端请求信息日志函数
func AccessLog(ip, method, url, userAgent string, stauts int) {
	log.Printf(`%s %s %s "%s" %d`, ip, method, url, userAgent, stauts)
}

// 通过闭包的手段
// 编写日志包装函数，等会在路由中调用该函数，因为我们的 GetUser AddUser DeleteUser Edit 等多个函数满足 http.HandlerFunc 方法
func LoggerWrapper(action http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		action(rw, r)
		AccessLog(r.RemoteAddr, r.Method, r.URL.String(), r.Header.Get("User-Agent"), 200)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("web").ParseFiles("template/user.html")
	if err != nil {
		log.Println(err)
		return
	}
	tpl.ExecuteTemplate(rw, "user.html", service.ListUser())
}

func AddUser(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.New("AddUser").ParseFiles("template/create.html")
		if err != nil {
			log.Println(err)
			return
		}
		tpl.ExecuteTemplate(rw, "create.html", nil)
	} else {
		service.AddUser(
			r.FormValue("name"),
			r.FormValue("age"),
			r.FormValue("addr"),
			r.FormValue("sex") == "1",
		)
	}
	http.Redirect(rw, r, "/", 302)

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
		service.DeleteUser(id)
	}
	http.Redirect(rw, r, "/", 302)
}

func Edit(rw http.ResponseWriter, r *http.Request) {
	type Newuser struct {
		ID   int64
		Name string
		Age  string
		Sex  bool
		Addr string
	}

	if r.Method == "GET" {
		tlp, err := template.New("edit").ParseFiles("template/edit.html")
		if err != nil {
			log.Println(err)
			return
		}
		r.ParseForm()
		fmt.Println("r.Form from /edit/:", r.Form.Get("id"))
		i := r.FormValue("id")
		id, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		user, err := service.IDFindUser(id)
		newuser := Newuser{
			ID:   user.ID,
			Name: user.Name,
			Age:  user.Age,
			Sex:  user.Sex,
			Addr: user.Addr,
		}
		fmt.Println("user id = ", user.ID)
		if err != nil {
			log.Println(err)
		}
		tlp.ExecuteTemplate(rw, "edit.html", newuser)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("from /edit/ r.PostFrom: ", r.Form.Get("id"))
		id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("edit user, id is: ", id)
		var (
			name = r.FormValue("name")
			age  = r.FormValue("age")
			addr = r.FormValue("addr")
			sex  = r.FormValue("sex") == "1"
		)
		if err := service.Edit(id, name, age, addr, sex); err != nil {
			fmt.Println(err)
			return
		}
	}
	http.Redirect(rw, r, "/", 302)
}
