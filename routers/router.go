package routers

import (
	"hmcms/controllers"
	"hmcms/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/admin/index", &admin.IndexController{}, "get:Index")
}
