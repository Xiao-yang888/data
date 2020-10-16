package models

import "data/db_mysql"

/*
 *上传文件的记录
 */
type upload struct {
	Id int
	UserId int
	FileName string
	FileSize int
	FileCert string
	FileTitle string
	CerTime int
}

func (u upload) SavaRecord() {
	db_mysql.Db.Exec("insert into ...")

}
