/*
go-pg
	1. 数据库的连接
	2. 数据表的创建
	3. SQL增删查改操作
	4. ...
*/
package main

import (
	_ "fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	mdb "play-pg/db"
)

func main() {
	option := mdb.IndividualLocalConnection()
	db := pg.Connect(&option)
	defer db.Close()

	if err := db.Ping(db.Context()); err != nil {
		log.Fatal("Ping error")
	}

	// pb := db.PolicyBasic{}
	// if err := db.Model(&pb).Where("id = ?", 85).Select(); err != nil {
	// 	log.Fatal("Select error")
	// }
	// fmt.Printf("%+v\n", pb)

	createTable(db)
}

func createTable(db *pg.DB) {
	err := db.Model((*mdb.Account)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal("Create table error:", err)
	}
}

