package models

type User struct {
	Id string `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}
