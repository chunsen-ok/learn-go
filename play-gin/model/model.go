package model

type RustPost struct {
	tableName struct{} `pg:"rust_post"`
	Id        int64    `pg:",pk"`
	Name      string
	Data      []byte
	Mdt       string
}
