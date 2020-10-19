package models

import (
	"data/db_mysql"
	"data/utils"
	"fmt"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

/**
 *将用户的信息保存到数据库中
 */
func (u User) AddUser() (int64,error) {
	//把脱密的密码的md5值重新赋值为密码进行储存
    u.Password = utils.MD5HashString(u.Password)
    //fmt.Println(u.Password,u.Phone)

    rs, err := db_mysql.Db.Exec("insert into data(phone,password) values(?,?)",
		u.Phone, u.Password)
	//错误早发现早解决
	if err != nil { //保存数据遇到错误
		fmt.Println(err.Error())
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	//id值代表的是此次数据操作影响的行数，id是一个整数int64类型
	return id, nil
}

/**
 *查询用户信息
 */
func (u User) QueryUser() (*User,error) {

	//把加密的密码的md5值重新赋值为密码进行储存
	u.Password = utils.MD5HashString(u.Password)
	row := db_mysql.Db.QueryRow("select phone from data where phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

func (u User) QueryUserByPhone() (*User, error) {
	fmt.Println(u.Phone)
	row := db_mysql.Db.QueryRow("select id from data where phone = ?",u.Phone)
	var user User
	err := row.Scan(&user.Phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

