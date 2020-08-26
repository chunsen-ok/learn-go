package db


import "github.com/go-pg/pg/v10"


func CorpConnection() pg.Options {
	return pg.Options{
		Addr:     "127.0.0.1:5432",
		User:     "bt26j749o4jiju3jv1l0",
		Password: "bt26j749o4jiju3jv1kg",
		Database: "bt26j749o4jiju3jv1l0",
	}
}

func IndividualLocalConnection() pg.Options {
	return pg.Options {
		Addr: "127.0.0.1:5432",
		User: "postgres",
		Password: "pg2020",
		Database: "postgres",
	}
}

func IndividualNetConnection() pg.Options {
	return pg.Options {
		Addr: "47.107.106.105:5432",
		User: "cx_rw",
		Password: "CX_RW_1397",
		Database: "cx_dev",
	}
}
