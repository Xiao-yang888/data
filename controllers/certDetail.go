package controllers

import (
	"data/blockchain"
	"data/models"
	"data/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type CertDetailController struct {
	beego.Controller
}

/**
 *该get方法用于接收处理浏览器的get请求，网查看证书详情页面跳转
 */
func (c* CertDetailController) Get() {
	//1，解析和接收前端页面传递的数据cert_id
	cert_id := c.GetString("cert_id")

	//2，到区块链上查询区块数据
	//fmt.Println("要查询的认证数据的id：", cert_id)
	block, err := blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if err != nil {
		c.Ctx.WriteString("抱歉，查询链上数据遇到错误，请重试!")
		return
	}
	if block == nil {//遍历整条区块链，但是未查到数据
		c.Ctx.WriteString("抱歉，未查询到链上数据")
		return
	}
	fmt.Println("查询到的区块的高度：", block.Height)

	//反序列化
	certRecord, err := models.DeserializeCertRecord(block.Data)
    certRecord.CertIdHex = strings.ToUpper(string(certRecord.CertId))
    certRecord.CertHashHex = string(certRecord.CertHash)
	certRecord.CertTimeFormat = utils.TimeFormat(certRecord.CertTime, utils.TIME_FORMAT_ONE)
	//结构体
	c.Data["CertRecord"] = certRecord
	
	//3，跳转证书详情页面
	c.TplName = "cert_detail.html"
}
