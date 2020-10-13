package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
//展示默认首页，注册页面
func (l *LoginController) Get() {
	l.TplName = "login.html"
}
