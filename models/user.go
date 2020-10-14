package models

import (
	"crypto/md5"
	"data/db_mysql"
	"encoding/hex"
)

type User struct {
	Id string `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

/**
 *将用户的信息保存到数据库中
 */
func (u User) AddUser() (int64,error) {
	//加密
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes) //把加密的密码的md5值虫嗪赋值为密码进行储存

	rs, err := db_mysql.Db.Exec("insert  into user (phone,password) values (?,?)",
		u.Phone, u.Password)
	//错误早发现早解决
	if err != nil {
		//保存数据遇到错误
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		//保存数据遇到错误
		return -1, err
	}
	//id值代表的是此次数据操作影响的行数，id是一个整数int64类型
	return id, nil
}

/**
 *查询用户信息
 */
func (u User) QueryUser() (*User,error) {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(u.Password))
	pwdBytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(pwdBytes) //把加密的密码的md5值虫嗪赋值为密码进行储存

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

