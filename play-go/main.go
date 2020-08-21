/*
基础 standard package：
	builtin
	fmt
*/
package main

import (
	_ "fmt"
	"github.com/LCSyy/play-go/play"
)

func main() {
	app := play.NewApp()
	app.Run()
}

