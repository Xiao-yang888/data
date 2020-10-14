package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//展示默认首页，注册页面
func (c *MainController) Get() {
	fmt.Println("hello world")
	c.TplName = "register.html"
}


