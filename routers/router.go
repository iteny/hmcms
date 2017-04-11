package routers

import (
	"github.com/astaxie/beego"
	"xhcms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
}
