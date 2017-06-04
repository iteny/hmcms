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
	fmt.Println(rows)
	c.TplName = "admin/site/menu.html"
}
