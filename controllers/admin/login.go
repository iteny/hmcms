package admin

import (
	// "github.com/astaxie/beego"

	"hmcms/models"
	"hmcms/models/sqlite"
	"log"
	"time"

	"github.com/astaxie/beego/cache"

	"github.com/astaxie/beego/validation"
)

type LoginController struct {
	BaseController
}
type Users struct {
	cnt int
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
		// fmt.Println(row)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": 4, "info": "发生未知错误！"}
			c.ServeJSON()
			return
		} else {
			s, err := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
			if err != nil {
				log.Println(s)
			}
			errorstatus := s.Get("error_" + username)

			log.Println(errorstatus)
			if errorstatus != nil {
				c.Data["json"] = map[string]interface{}{"status": 4, "info": "帐号密码输入错误3次，请在10分钟以后在登录！"}
				c.ServeJSON()
				return
			}
			if row.Id != 0 {
				ip := c.Ctx.Input.IP()
				ipinfo := models.TaobaoIP(ip)
				if row.Status == 0 {
					c.Data["json"] = map[string]interface{}{"status": 4, "info": "帐号已被禁用！"}
					c.ServeJSON()
					return
				} else if row.Status == 1 {
					sqlite.LoginLogDb.RecordLogin(row.Username, ip, 1, "登录成功", ipinfo.Data.Area, ipinfo.Data.Country, c.Ctx.Input.UserAgent())
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
				var status int
				if Users.cnt > 3 {
					status = 1
				}
				Users.cnt++

				s.Put("error_"+username, status, 10*time.Second)
				c.Data["json"] = map[string]interface{}{"status": 4, "info": "用户名或密码错误！"}
				c.ServeJSON()
				return
			}
		}

	}

}
