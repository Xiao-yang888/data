package routers

import (
	"data/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/register.html",&controllers.RegisterController{})
    //用户登陆接口
    beego.Router("/login",&controllers.LoginController{})

    beego.Router("/login.html",&controllers.LoginController{})

    beego.Router("/list_record.html", &controllers.UploadFileController{})

    beego.Router("/UploadFile",&controllers.UploadFileController{})

    //查看认证数据证书
    beego.Router("/cert_detail.html",&controllers.CertDetailController{})

    //用户实名认证
    beego.Router("user_kyc",&controllers.UserKycController{})
}
