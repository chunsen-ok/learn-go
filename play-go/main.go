/*
go standard package
	builtin V
	fmt V
	errors V
	strings
	strconv
	math
	time
*/
package main

import (
	"fmt"
	"strconv"
)

type MetaObject struct {
	Id   int
	Name string
}

type Object struct {
	MetaObject
}

func main() {
	objMap := make(map[int]Object)
	for i := 0; i < 10; i = i + 1 {
		obj, ok := objMap[i]
		if !ok {
			obj = Object{
				MetaObject: MetaObject{
					Id:   i,
					Name: strconv.Itoa(i),
				},
			}
		}
		obj.Name = "Update" + obj.Name
		objMap[i] = obj
	}

	for id, obj := range objMap {
		fmt.Println(id, ":", obj.Name)
	}
}
