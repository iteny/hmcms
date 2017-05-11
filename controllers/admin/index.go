package admin

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	c.VerifyLogin()
	c.Data["username"] = c.GetSession("username")
	c.TplName = "admin/index.html"
}
