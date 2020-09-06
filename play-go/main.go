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
	"strings"
)

func main() {
	s := "abcabc"
	ss := "abc"
	fmt.Println(strings.Count(s,ss))
	fmt.Println(strings.EqualFold("Ab","aB"))
	fmt.Println(strings.Fields("This is a word"))
}

