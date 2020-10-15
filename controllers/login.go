package controllers

import (
	"data/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//展示默认首页，注册页面
func (l *LoginController) Get() {
	l.TplName = "login.html"
}


/**
 *post方法处理用户的登录请求
 **/

func (l *LoginController) Post() {
	//解析客户端用户提交的登陆数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登陆信息解析失败，请重试！")
		return
	}

	//根据解析到的数据，执行数据库查询操作
	u, err  := user.QueryUser()

	//判断数据库查询结果
	if err != nil {
		//sql : no rows in result set(集合)，结果集中没有数据
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登陆失败，请重试！")
		return
	}

	//根据查询结果返回客户端相应的信息或页面跳转
	l.Data["phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"//文件上传界面
}