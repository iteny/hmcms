package admin

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"database/sql"
	"github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type LoginController struct {
	beego.Controller
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
	username := c.GetString("username")
	password := c.GetString("password")
	valid := validation.Validation{}
	valid.Required(username, "请输入用户名")
	valid.MinSize(username, 6, "用户名最少6位")
	valid.MaxSize(username, 15, "用户名最多15位")
	valid.Required(password, "请输入密码")
	valid.MinSize(password, 8, "密码最少8位")
	valid.MaxSize(password, 15, "密码最多15位")
	// // valid.Range(u.Age, 0, 18, "age")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			// log.Println(err.Key, err.Message)
			// log.Println(username, password)
			c.Data["json"] = map[string]interface{}{"status": 4, "info": err.Key}
			c.ServeJSON()
		}
	} else {
		db, err := sql.Open("sqlite3", "./sql/hmcms.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		rows, err := db.Query("select * from user")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []Users = make([]Users, 0)
		for rows.Next() {
			var u Users
			rows.Scan(&u.username, &u.password)
			users = append(users, u)
		}
		log.Println(users)
		c.Data["json"] = map[string]interface{}{"status": 4, "info": "ok"}
		c.ServeJSON()

	}
	// log.Println(username, password, 111)
	// // c.AllowCross() //允许跨域
	// c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	// c.ServeJSON()
}
