package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	block0 := blockchain.CreateGenesisBlock()//创建创世区块
	//fmt.Println(block0)
	block1 := blockchain.NewBlock(
		block0.Height + 1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block的哈希：%x\n", block0.Hash)
	fmt.Printf("block1的哈希: %x\n", block1.Hash)
	fmt.Printf("block1的PrevHash: %x\n", block1.PrevHash)

	//序列化
	blockJson, _ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block：", string(blockJson))
	return

	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()


}

