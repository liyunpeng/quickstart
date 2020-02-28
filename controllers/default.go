package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//@router /backlogin/judge [post]
func (this *MainController)JudgeLogin()  {
	this.Data["json"] = map[string]interface{}{"name": 0, "message": "你输入的账号或密码不正确"}
	this.ServeJSON()
}