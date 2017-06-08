package admin

import (
	"encoding/json"
	"fmt"
	"hmcms/models/sqlite"
	"sort"
)

type SiteController struct {
	BaseController
}
type Menu struct {
	Id   int `json:"id"`
	Sort int `json:"sort"`
}

//菜单列表
func (c *SiteController) Menu() {
	rows, _ := sqlite.AuthRuleDb.GetAllMenu()
	ar := sqlite.RecursiveMenu(rows, 0, 0)
	// fmt.Println(rows)
	// fmt.Println(ar)
	c.Data["json"] = ar
	c.TplName = "admin/site/menu.html"
}

//菜单排序
func (c *SiteController) SortMenu() {
	ob := make(map[int]string, 0)
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	var obk []int
	// obv := make(map[int]string, 0)
	for k, _ := range ob {
		obk = append(obk, k)
		// obv[k] = v
	}
	sort.Ints(obk)
	// var obs []int
	// for k, _ := range obv {
	// 	obs = append(obs, k)
	// }
	fmt.Println("obk=", obk)
	// fmt.Println("obv=", obv)
	// var buffer bytes.Buffer
	// for _, v := range ob {
	// 	buffer.WriteString(v)
	// 	buffer.WriteString(",")
	// }
	// fmt.Println(buffer.String())
	c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误！"}
	c.ServeJSON()
}
