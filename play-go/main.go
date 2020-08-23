/*
基础 standard package：
	builtin
	fmt
*/
package main

import (
	"fmt"
	_ "github.com/LCSyy/play-go/play"
)

type One struct {
	tableName struct {} `pg:"TableName"`
	id int64
	name string
}

func main() {
	// app := play.NewApp()
	// app.Run()
	o := One {
		id: 12314324,
		name: "sfsdf",
	}
	fmt.Printf("%+v",o)
}
