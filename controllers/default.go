package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//展示默认首页，注册页面
func (c *MainController) Get() {
	c.TplName = "register.html"
}


