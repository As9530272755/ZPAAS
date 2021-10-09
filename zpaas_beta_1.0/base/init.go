package base

import (
	"log"

	"zpaas_beta_1.0/models"
	"zpaas_beta_1.0/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type RequiredAuthController struct {
	beego.Controller
}

// 初始化 session
func (c *RequiredAuthController) Prepare() {
	c.Data["currentUser"] = nil
	c.Data["navKey"] = ""
	// 通过 session 获取 id 传入 user 变量
	user := c.GetSession("user")

	// 判断用户是否存在，如果为 nil 跳转到 login 登录页面
	if user == nil {
		c.Redirect("/login", 302)
		return
	}

	// 如果用户存在通过断言判断是否是 int64
	if pk, ok := user.(int64); ok {
		// 传入 id 到 service.GetById() 获取到对应的 user 结构体
		if user := service.GetById(pk); user != nil {
			// 然后再将对应的结构体传给 c.Data
			c.Data["currentUser"] = user
		}
	}

	// 如果现在用户未登录且有了 seesion 也直接销毁 session 并且跳转之 login
	if c.Data["currentUser"] == nil {
		c.DestroySession()
		c.Redirect("/login", 302)
		return
	}

}

// 初始化数据库
func DbRegister() {
	user := beego.AppConfig.String("db::Mysqluser")
	password := beego.AppConfig.String("db::Mysqlpassword")
	host := beego.AppConfig.String("db::Mysqlhost")
	Port := beego.AppConfig.String("db::MysqlPort")
	db := beego.AppConfig.String("db::MysqlDataBase")
	dsn := user + ":" + password + "@tcp" + "(" + host + ":" + Port + ")" + "/" + db + "?charset=utf8mb4&loc=Local&parseTime=true"

	// 注册数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(&models.User{})
	orm.RunSyncdb("default", false, true)
	beego.Informational("注册数据库")
	InitAdmin()
}

// 初始化管理员
func InitAdmin() {
	if user := service.GetUserByName("root"); user == nil {
		log.Println("创建管理员")
		newpassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), 5)
		adminUser := models.User{
			Name:     "root",
			Password: string(newpassword),
		}
		service.Add(&adminUser)
	}
	beego.Informational("创建初始化管理员")
}
