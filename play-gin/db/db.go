package db

import (
	"fmt"
	_ "log"
	"github.com/go-pg/pg/v10"
)

type RustPost struct {
	tableName struct {} `pg:"rust_post"`
	Id int64
	Name string
	Data []byte
	Mdt string
}

func (rp RustPost) String() string {
	return fmt.Sprintf("%+v",rp)
}

func Setup() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr: "47.107.106.105:5432",
		User: "cx_rw",
		Password: "CX_RW_1397",
		Database: "cx_dev",
	})

	if err := db.Ping(db.Context()); err != nil {
		fmt.Println("Error:", err)
	}

	return db
}


func Query(db *pg.DB, id int64) string {
	rustPost := RustPost {}
	err := db.Model(&rustPost).
		Where("id = ?", id).
		Select()
	if err != nil {
		fmt.Println("Error:",err)
	}

	fmt.Println("id: ",rustPost.Id)
	fmt.Println("name: ",rustPost.Name)
	fmt.Println("mdt: ", rustPost.Mdt)
	
	return rustPost.Name
}
