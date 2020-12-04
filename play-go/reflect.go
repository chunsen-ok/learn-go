package main

import (
	"fmt"
	"reflect"
)

type Everything struct {
	Id   int
	Name string
}

func reflectPlayground() {
	v := Everything{}
	t := reflect.TypeOf(v)

	fmt.Println("name:", t.Name())
	fmt.Println("pkg path:", t.PkgPath())
	fmt.Println("string:", t.String())
	fmt.Println("align:", t.Align())
	fmt.Println("field align:", t.FieldAlign())
	fmt.Println("num of field:", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name, ":", field.Type.Name())
	}
}
