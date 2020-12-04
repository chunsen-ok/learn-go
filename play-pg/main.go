/*
go-pg
	1. 数据库的连接
	2. 数据表的创建
	3. SQL增删查改操作
	4. ...
*/
package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type MyUser struct {
	Id       int64
	What     *float64
	Name     string
	Account  string
	Password string
	Comment  string
	Protect  []Protect
	Action   Action
}

type Protect struct {
	Subject string
	Answer  string
}

type Action struct {
	CreateTime string
	ModifyTime []string
}

func main() {
	// options := pg.Options{
	// 	Addr:     "127.0.0.1:5432",
	// 	User:     "postgres",
	// 	Password: "pg2020",
	// 	Database: "postgres",
	// }
	// db := pg.Connect(&options)
	// defer db.Close()

	// if err := db.Ping(db.Context()); err != nil {
	// 	log.Println("Can't ping db: ", err)
	// }

	// createTable(db)

	/*
		data := MyUser{
			Id:       1,
			Name:     "lcs",
			Account:  "chuns.liu@foxmail.com",
			Password: "adfasdfadfsasdf",
		}

		_, err := db.Model(&data).Insert()
		if err != nil {
			log.Println("Insert my_user error: ", err)
		}
	*/

	// insertData(db)
}

func createTable(db *pg.DB) {
	ctOptions := orm.CreateTableOptions{
		IfNotExists: true,
	}
	if err := db.Model((*MyUser)(nil)).CreateTable(&ctOptions); err != nil {
		log.Println("Create table error: ", err)
	}

}

func insertData(db *pg.DB) {
	data := MyUser{
		Id:       1,
		Name:     "lcs",
		Account:  "chuns.liu@foxmail.com",
		Password: "adfasdfadfsasdf",
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("Begin Tx error: ", err)
	}

	count, err := tx.Model((*MyUser)(nil)).Count()
	if err != nil {
		log.Println("Select count error: ", err)
		tx.Rollback()
	}

	if count > 0 {
		fmt.Println("Count is: ", count)
	}

	_, err = tx.Model(&data).Column("password").WherePK().Update()
	if err != nil {
		log.Println("Insert data error: ", err)
		tx.Rollback()
		return
	}

	err = tx.Model(&data).WherePK().Select()
	if err != nil {
		log.Println("Select data error: ", err)
		tx.Rollback()
		return
	}

	tx.Commit()

}
