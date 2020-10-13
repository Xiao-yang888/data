package routers

import (
	"data/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/register.html",&controllers.RegisterController{})

    beego.Router("/login",&controllers.LoginController{})

    beego.Router("/login.html",&controllers.LoginController{})

}
