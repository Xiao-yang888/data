package controllers

import (
	"data/db_mysql"
	"data/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
//展示默认首页，注册页面
func (r *RegisterController) Get() {

	var user models.User
	err:= r.ParseForm(user)
	if err != nil {
		return
	}


	row, err := db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败，请重试！")
		return
	}
	fmt.Println(row)
	//row代表影响了几行

	if row != -1 {
		r.TplName = "login.html"
	}else{
		r.TplName = "error.html"
	}
}
