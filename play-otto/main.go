package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/robertkrimen/otto"
)

type Me struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MeInfo struct {
	Me
	Age string `json:"age"`
}

func main() {
	me := MeInfo{
		Me: Me{
			Id:   54235,
			Name: "sdfdsfs",
		},
		Age: "fsdf",
	}
	b, err := json.Marshal(me)
	if err != nil {
		fmt.Println("marshal error:", err)
	}
	fmt.Println(string(b))

	vm := otto.New()
	obj, err := vm.Object("(" + string(b) + ")")
	if err != nil {
		log.Fatal("object error:", err)
	}
	vm.Set("gg", obj)
	vm.Run("abc = 2 + 2;console.log('value is:',JSON.stringify(gg));console.log('This value is:',abc);")

	fmt.Println(me.Me.Id)
}
