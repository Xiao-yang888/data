package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"github.com/astaxie/beego"
)

func main() {

	//先准备一条区块链
	blockchain.NewBlockChain()

	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()


}

