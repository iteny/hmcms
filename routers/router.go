package routers

import (
	"github.com/astaxie/beego"
	"hmcms/controllers"
	"hmcms/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
	// beego.Router("/admin/login", &admin.LoginController{}, "post:Login")
}
