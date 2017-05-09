package admin

import (
	// "github.com/astaxie/beego"

	"fmt"
	"hmcms/models/sqlite"

	"github.com/astaxie/beego/validation"
	"github.com/yinheli/qqwry"
)

type LoginController struct {
	BaseController
}
type Users struct {
	username int
	password string
}

func (c *LoginController) GetLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/login.html"

}
func (c *LoginController) PostLogin() {
	var username string = c.GetString("username")
	password := c.GetString("password")
	valid := validation.Validation{}
	valid.Required(username, "请输入用户名")
	valid.MinSize(username, 5, "用户名最少5位")
	valid.MaxSize(username, 15, "用户名最多15位")
	valid.Required(password, "请输入密码")
	valid.MinSize(password, 8, "密码最少8位")
	valid.MaxSize(password, 15, "密码最多15位")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			c.Data["json"] = map[string]interface{}{"status": 4, "info": err.Key}
			c.ServeJSON()
			return
		}
	} else {

		row, err := sqlite.UserDb.LoginUser(username, password)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误！"}
			c.ServeJSON()
			return
		} else {
			if row != nil {
				if row.Status == 0 {

					c.Data["json"] = map[string]interface{}{"status": 4, "info": "帐号已被禁用！"}
					c.ServeJSON()
					return
				} else if row.Status == 1 {

					ip := "180.165.111.135"
					// resp, err := http.Get(fmt.Sprintf("http://ip.taobao.com/service/getIpInfo.php?ip=%s", ip))
					// if err != nil {
					// 	fmt.Println("获取失败")
					// }
					// defer resp.Body.Close()
					// fmt.Println(resp)
					file := "./ipdata/UTFWry.dat"
					q := qqwry.NewQQwry(file)
					q.Find(ip)
					fmt.Printf("ip:%v, Country:%v, City:%v", q.Ip, q.Country, q.City)
					ss := c.Ctx.Input.IP()
					fmt.Println(ss)
					c.SetSession("userid", row.Id)
					c.SetSession("username", row.Username)
					c.SetSession("status", row.Status)

					c.Data["json"] = map[string]interface{}{"status": 1, "info": "登录成功！"}
					c.ServeJSON()
					return
				} else {
					c.Data["json"] = map[string]interface{}{"status": 4, "info": "未知错误！"}
					c.ServeJSON()
					return
				}

			} else {
				c.Data["json"] = map[string]interface{}{"status": 4, "info": "用户名或密码错误！"}
				c.ServeJSON()
				return
			}
		}

		// a, err := sqlite.GetAccount(1)
		// if err != nil {
		// 	log.Println(err)
		// } else {
		// 	log.Fatalf("%#v\n", a)
		// }

	}
	// log.Println(username, password, 111)
	// // c.AllowCross() //允许跨域
	// c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	// c.ServeJSON()
}
