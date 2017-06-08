package routers

import (
	"hmcms/controllers"
	"hmcms/controllers/admin"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

func init() {
	beego.SetStaticPath("/static", "static")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.LoginController{}, "get:GetLogin;post:PostLogin")
	beego.Router("/admin/index", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/index/home", &admin.IndexController{}, "get:Home")
	beego.Router("/admin/getLeftMenu", &admin.IndexController{}, "post:GetLeftMenu")
	beego.Handler("/captcha/*.png", captcha.Server(150, 36)) //注册验证码服务，验证码图片的宽高为240 x 80
	beego.Router("/admin/site/menu", &admin.SiteController{}, "get:Menu")
	beego.Router("/admin/site/sortmenu", &admin.SiteController{}, "post:SortMenu")
}
