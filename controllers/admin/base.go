package admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) VerifyLogin() {
	userId := c.GetSession("userId")
	if userId == nil {
		c.Redirect("/admin", 302)
	}
}
func (c *BaseController) RecordLogin(status int, message string) {

}
