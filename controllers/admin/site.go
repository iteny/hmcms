package admin

import (
	"fmt"
	"hmcms/models/sqlite"
)

type SiteController struct {
	BaseController
}

func (c *SiteController) Menu() {
	rows, _ := sqlite.AuthRuleDb.GetAllMenu()
	ar := sqlite.RecursiveMenu(rows, 0, 0)
	// fmt.Println(rows)
	fmt.Println(ar)
	c.Data["json"] = ar
	c.TplName = "admin/site/menu.html"
}
