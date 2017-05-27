package admin

import (
	"fmt"
	"hmcms/models/sqlite"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	rows, _ := sqlite.AuthRuleDb.GetOneMenu()
	fmt.Println(rows, 111)
	c.Data["menu"] = rows
	c.VerifyLogin()
	c.Data["username"] = c.GetSession("username")
	c.TplName = "admin/index.html"
}

//获取顶部菜单
// func (c *IndexController) GetOneMenu() {
//
// }
