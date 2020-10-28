package main

import (
	"data/blockchain"
	"data/db_mysql"
	_ "data/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {

	//block0 := blockchain.CreateGenesisBlock()//创建创世区块
	//block1 := blockchain.NewBlock(
	//	block0.Height + 1,
	//	block0.Hash,
	//	[]byte{})
	//fmt.Printf("block的哈希：%x\n", block0.Hash)
	//fmt.Printf("block1的哈希: %x\n", block1.Hash)
	//fmt.Printf("block1的PrevHash: %x\n", block1.PrevHash)
	//
	//
	//block0Bytes := block0.Serialize()
	//fmt.Println("创世区块gob序列化后：",block0Bytes)
	//deblock0, err := blockchain.DeSerialize(block0Bytes)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("反序列化后区块的高度：", deblock0.Height)
	//fmt.Println("反序列化后的区块的哈希", deblock0.Hash)

	bc := blockchain.NewBlockChain()//封装
	fmt.Printf("创世区块的hash：%x\n", bc.LastHash)
	//bc.AddData([]byte("用户的数据"))
	return

	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映射
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()


}

