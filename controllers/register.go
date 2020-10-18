package controllers

import (
	"data/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//展示默认首页，注册页面
func (r *RegisterController) Post() {

	var user models.User
	err:= r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重试！")
		return
	}

	_, err = user.AddUser()
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，用户注册失败，请重试！")
		return
	}

	r.TplName = "login.html"

    ////将解析到的数据保存到数据库中
	//row, err := db_mysql.AddUser(user)
	//if err != nil {
	//	r.Ctx.WriteString("注册用户信息失败，请重试！")
	//	return
	//}

	//row代表影响了几行

	//if row != -1 {
	//	r.TplName = "login.html"
	//}else{
	//	r.TplName = "error.html"
	//}
}
