package main

import (
	"fmt"
	"github.com/LCSyy/play-go/play"
)

func main() {
	obj := play.NewObject("Monster")
	fmt.Println("The demon has a name:", obj.ObjectName)
	fmt.Println("The demon is in ", obj.Status())
	if obj.Status() == play.Dead {
		fmt.Println("The demon is dead!")
	}

	objPtr := new(play.Object)
	fmt.Println(objPtr.ObjectName, " ", objPtr.Status())

	obj2 := play.Object{ ObjectName: "No name" }
	fmt.Println(obj2.ObjectName)

	var arr = [12]int{1,44,57,601,73,84,925,12,546,87,34,212}
	for key, val := range arr {
		min,max := GetMultiResult(val)
		defer ShowIndex(key)
		fmt.Println("min:",min, " max:", max)
	}
}

func GetMultiResult(val int) (int,int) {
	times := val / 10
	// remainder := val % 10
	return times * 10, (times + 1) * 10
}

func ShowIndex(idx int) {
	fmt.Println("index:", idx)
}
