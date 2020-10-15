package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
 *该控制器结构体用于处理文件上传的功能
 */
type UploadFileController struct {
	beego.Controller
}

/**
 *该post方法用于处理用户在客户端提交的文件
 */
func (u *UploadFileController) Post() {
	title := u.Ctx.Request.PostFormValue("upload_title")//用户输入的标题
	fmt.Println("电子数据标签:",title)
	file, header, err := u.GetFile("zengyang")
	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}
	defer file.Close()//延迟执行 空指针错误：invalid memorey or nil pointer derefernce

	//使用io包提供的方法保存文件
	saveFilePath := "static/upload/" + header.Filename
	saveFile, err :=os.OpenFile(saveFilePath,os.O_CREATE || os.O_RDWR,777)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
		return
	}
	_, err = io.Copy(saveFile,file)
	if err != nil{
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试！")
		return
	}

	//计算文件的SHA256值
	hash256 := sha256.New()
	fileBytes, _ := ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes := hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))

    //先查询用户ID

    //把上传的文件作为记录保存到数据库中

	u.Ctx.WriteString("恭喜，已接收到上传文件")
}


/**
 *该post方法用户处理用户在客户端提交的认证文件
 */
func (u *UploadFileController) Post1() {
	//解析用户上传的数据及内容

	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title")//用户输入的标题

	//用户上传的文件
	file, header, err := u.GetFile("zengyang")
	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试！")
		return
	}

	defer file.Close()

	fmt.Println("自定义的标题：",title)
	//获得到了上传的文件
	fmt.Println("上传的文件名称：",header.Filename)
	//eg: 支持jpg，png类型，不支持jpeg，gif类型
	//文件名：文件名 + "." + 扩展名
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(fileNameSlice)
	fmt.Println(":",strings.TrimSpace(fileType))
	isJpg := strings.HasSuffix(header.Filename,"jpg")
	isPng := strings.HasSuffix(header.Filename,"png")
	if  !isJpg && isPng {
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传符合格式的文件")
		return
	}

	//文件的大小 200kb
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")

	if header.Size / 1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传符合要求的文件")
		return
	}

	fmt.Println("上传文件的大小：",header.Size)//字节大小


	//构造文件名称
	//rand,Seed(time.Now().UnixNano())
	//randNum := fmt.Sprintf("%d",rand,Intn(9999)+1000)
	//hashName := md5.Sum([]byte( time.Now().Format("2006_01_02_15_04_05")+

	//fileName := fmt.Sprintf("%x", hashName) +ext
	//this.Ctx.WriteString(fileName)

	//perm:permission 权限
	//权限的组成： a+b+c
	    //a:文件所有者对文件的操作权限    读4，写2，执行1
	    //b：文件所有者所在组的用户对文件的操作权限   读4，写2，执行1
	    //c：其他用户对文件的操作权限   读4，写2，执行1

	    //eg:m文件，权限是651。
	    //判断题：文件所有者对该m文件有写权限（dui）
	 saveDir := "static/upload"
	 //尝试打开文件夹
	 _, err = os.Open(saveDir)
	 //os.OpenFile(name文件名,flag文件的操作项，例如os.CREATE,OS.O_RDWR etc,perm权限)
	 if err != nil{//打开失败:文件夹不存在
	 	//创建文件夹
		err = os.Mkdir(saveDir,777)
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试！")
			return
		}
	}


	//if err != nil {
	// 	fmt.Println(err.Error())
	// 	//打开文件目录遇到错误
	// 	u.Ctx.WriteString("打开文件失败")
	//	 return
	// }
	// fmt.Println("打开的文件夹:",f.Name())

	 //文件名： 文件路径 + 文件名 + "." +文件扩展名
	 saveName := "static/upload/" + header.Filename
	 fmt.Println("要保存的文件名：",saveName)

	//fromFile:文件，
	//toFile:要保存的文件路径
	err = u.SaveToFile("zengyang",saveName)
	 if err != nil {
	 	fmt.Println(err.Error())
	 	u.Ctx.WriteString("抱歉，文件认证失败，请重试！")
		 return
	 }

	fmt.Println("上传的文件：",file)

	u.Ctx.WriteString("已获取到上传文件。")

}
