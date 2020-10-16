package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/mysql"
)
//数据库连接
var Db *sql.DB
func Connect(){
	fmt.Println("连接数据库")
    
	//项目配置
	config := beego.AppConfig//定义config变量，接收并赋值为全局变量
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPwd := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	//db.Open连接数据库，有两个参数
	db,err := sql.Open(dbDriver,dbUser+":"+dbPwd+"@tpc("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil{//err 不为nil时 连接数据库发生错误，程序就在此中断，无需再向下执行
		//早发现，早解决
		panic("数据库连接错误，请查找bug后再试")
		//使程序进入恐慌状态，程序会终止执行
	}
	Db = db
	//fmt.Println(db)
}
