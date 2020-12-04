// pgx
// module: github.com/jackc/pgx/v4

package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Financial struct {
	Name     string
	TaxNum   string
	Bank     string
	BankNum  string
	Addr     []string
	Tel      []string
	PostCode []string
	Fax      []string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:pg2020@localhost:5432/corp-erp")
	defer conn.Close(context.Background())

	f := Financial{}
	err = conn.QueryRow(context.Background(), "SELECT financial FROM customer WHERE id = 3").Scan(&f)
	if err != nil {
		fmt.Println("[error] ", err)
		return
	}

	_, err = conn.Exec(context.Background(), "UPDATE customer SET financial = $1 WHERE id = 1", f)
	if err != nil {
		fmt.Println("[error] ", err)
		return
	}

	fmt.Println(f)
}
