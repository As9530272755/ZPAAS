package router

import (
	"zpaas_beta_1.0/controller"
	"zpaas_beta_1.0/filters"

	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register() {
	// 插入 BeforeExec 和 AfterExec 过滤器在所有 url 执行统计
	beego.InsertFilter("/*", beego.BeforeExec, filters.BeferExec)

	// 在执行 controller 之前会进入到 BeferExec ， 在执行之后会进入到 AfterExec, false 表示执行到 AfterExec
	beego.InsertFilter("/*", beego.AfterExec, filters.AfterExec, false)

	// 注册路由
	beego.ErrorController(&controller.ErrorController{})

	// Prometheus 数据采集函数
	beego.Handler("/metrics/", promhttp.Handler())

	beego.Router("/", &controller.RootPath{}, "get:Root")
	beego.Router("/user", &controller.Getuser{}, "get:List")
	beego.Router("/user/add/", &controller.Adduser{}, "get:Add;post:Add")
	beego.Router("/user/delete", &controller.DeleteUser{}, "get:Delete")
	beego.Router("/user/edit", &controller.EditUser{}, "get:Edit;post:Edit")
	beego.Router("/login", &controller.Login{}, "get:Login;post:Login")
	beego.Router("/out", &controller.Outlogin{}, "get:Out")
	beego.Router("/webshell", &controller.WebShell{}, "get:WebShell;post:WebShell")
	beego.Router("/paas", &controller.Paas{}, "get:Paas;post:Paas")
	beego.Router("/prometheus", &controller.Prom{}, "get:Prom;post:Prom")
	beego.Router("/userjson", &controller.GetJson{}, "get:UserJson")
}
