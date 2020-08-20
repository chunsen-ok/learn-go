package main

import (
	"fmt"
	"github.com/LCSyy/play-go/play"
)

type Greet interface {
	Greeting() string
}

type First struct {
	name string
}

func main() {
	obj := play.NewObject("Monster")
	fmt.Println("The demon has a name:", obj.ObjectName)
	fmt.Println("The demon is in ", obj.Status())

	if obj.Status() == play.Dead {
		fmt.Println("The demon is dead!")
	}

	var f Greet
	f = First{name: "Grey"}

	switch f := f.(type) {
	case First:
		fmt.Println("Yes:" + f.Greeting())
		break
	}
}

func (self First) Greeting() string {
	return "Hello, I'm " + self.name + "!"
}
