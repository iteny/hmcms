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
	beego.Handler("/captcha/*.png", captcha.Server(150, 36)) //注册验证码服务，验证码图片的宽高为240 x 80
}
