package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {


	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()

	block0 := blockchain.CreateGenesisBlock()
	block1 := blockchain.NewBlock(block0.Height+1, block0.Hash, []byte("a"))
	fmt.Println(block0,block1)
	return
}

