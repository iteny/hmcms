package admin

import (
	"fmt"
	"hmcms/models/sqlite"
	"strconv"
)

type IndexController struct {
	BaseController
}

//主页面
func (c *IndexController) Index() {
	rows, _ := sqlite.AuthRuleDb.GetOneMenu()
	fmt.Println(rows, 111)
	c.Data["menu"] = rows
	// c.VerifyLogin()
	c.Data["username"] = c.GetSession("username")
	c.TplName = "admin/index.html"
}

//主页面home页面
func (c *IndexController) Home() {
	c.TplName = "admin/home.html"
}

//获取左侧菜单
func (c *IndexController) GetLeftMenu() {
	// pid := c.GetString("pid")
	pid := c.Input().Get("pid")
	intpid, _ := strconv.Atoi(pid)
	rows, _ := sqlite.AuthRuleDb.GetTwoMenu(intpid)

	for k, v := range rows {

		throws, _ := sqlite.AuthRuleDb.GetTwoMenu(v.Id)
		for tk, _ := range throws {
			rows[k].Children = append(rows[k].Children, throws[tk])
		}

		// for _, d := range throws {
		// 	node := sqlite.AuthRule{Id: v.Id, Title: v.Title, Children: sqlite.AuthRule{Id: d.Id, Title: d.Title}}
		// 	// sqlite.AuthRule.Children = append(sqlite.AuthRule.Children, &node)
		// }

	}
	// fmt.Println(intpid)
	fmt.Println(rows)
	c.Data["json"] = rows
	c.ServeJSON()
	return
}
