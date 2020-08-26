package db


type User struct {
	Id int64
	Name string
	Nickname string
	UserDetails UserInfo
	Details DetailInfo
}

type UserInfo struct {
	Account string
	AuthType string
}

type DetailInfo struct {
	CreateDt string
}


///////////////////////

type Account struct {
	Id int64
	Account string
	Passwd string
	email string
	phone string
	Descriotion string
}
