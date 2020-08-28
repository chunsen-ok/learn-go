/*
基础 standard package：
	builtin
	fmt
*/
package main

import (
	"fmt"
	"time"
)

type MyTime time.Time

func (self MyTime) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func main() {
	var name string = fmt.Sprintf("%%%v%%", "This is name")

	if len(name) > 0 {
		fmt.Println(name)
	}

	if name != "" {
		fmt.Println(name)
	}

	filters := []string{"s", "3"}

	fmt.Println(filters)
}
