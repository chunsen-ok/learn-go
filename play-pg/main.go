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
	_ "fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	mdb "play-pg/db"
)

func main() {
	option := mdb.IndividualLocalConnection()
	db := pg.Connect(&option)
	defer func() {
		fmt.Println("Defer call")
		db.Close()
	}()

	if err := db.Ping(db.Context()); err != nil {
		log.Fatal("")
	}

	if err := createTable(db); err != nil {
		log.Fatal("What:", err)
	}

	// if data, err := query(db); err != nil {
	// 	goto Finish
	// } else {
	// 	fmt.Println(data)
	// }

	lst := make([]mdb.Account, 0)

	lst = append(lst, mdb.Account{
		Id:      23434,
		Account: "nobody",
	})
	lst = append(lst, mdb.Account{
		Id:      354343,
		Account: "somebody",
	})

	Jobs(&lst)
}

func createTable(db *pg.DB) error {
	option := orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.Model((*mdb.Account)(nil)).CreateTable(&option)
	if err != nil {
		return err
	}

	err = db.Model((*mdb.CorpCertificate)(nil)).CreateTable(&option)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user_info(
		id serial not null,
		name varchar(255) not null,
		age int4
	);
	COMMENT ON COLUMN public.user_info.name IS '用户名称';`)
	if err != nil {
		return err
	}

	return nil
}

func query(db *pg.DB) (mdb.Account, error) {
	model := mdb.Account{}
	_, err := db.QueryOne(&model, "SELECT id FROM accounts;")
	return model, err
}

func Jobs(lst *[]mdb.Account) {
	for _, val := range *lst {
		fmt.Println("Value: ", val)
	}
}
