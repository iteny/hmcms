package admin

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/validation"
	// "hmcms/models"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) GetLogin() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin/login.html"
}
func (c *LoginController) PostLogin() {
	// var ob models.Object
	// json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	// objectid := models.AddOne(ob)
	// this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	// this.ServeJson()
	username := c.GetString("username")
	password := c.GetString("password")
	// valid := validation.Validation{}
	// valid.Required(username, "1")
	// valid.MinSize(username, 6, "2")
	// valid.MaxSize(username, 15, "3")
	// valid.Required(password, "1")
	// valid.MinSize(password, 6, "2")
	// valid.MaxSize(password, 15, "3")
	// // valid.Range(u.Age, 0, 18, "age")

	// if valid.HasErrors() {
	// 	// 如果有错误信息，证明验证没通过
	// 	// 打印错误信息
	// 	for _, err := range valid.Errors {
	// 		log.Println(err.Key, err.Message)
	// 		log.Println(username, password)
	// 	}
	// }
	log.Println(username, password, 111)
	// c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}
