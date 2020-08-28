package db

import (
	"time"
)

type User struct {
	Id          int64
	Name        string
	Nickname    string
	UserDetails UserInfo
	Details     DetailInfo
}

type UserInfo struct {
	Account  string
	AuthType string
}

type DetailInfo struct {
	CreateDt string
}

///////////////////////

type Account struct {
	Id          int
	Account     string
	Passwd      string
	email       string
	phone       string
	Descriotion string
}

type CorpCertificate struct {
	tableName      struct{} `pg:"corp_certificate"`
	Id             int
	Seq            int32        // 序号
	Certificate    string       // 资质证照
	No             string       // 文件编号
	IssueDate      time.Time    // 发证日期
	ExpirationDate time.Time    // 有效日期
	Remark         string       // 备注
	Attachment     []Attachment // 附件
}

type Attachment struct {
	Name string
	Md5  string
	Path string
}
