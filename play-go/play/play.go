/*
play 用于学习的包。
*/
package play

import (
	"fmt"
	"strconv"
)

const (
	EVAL = iota
	FETCH
	EXIT
)

var (
	appChan chan int
)

///////////////////////
// ScanArgs
///////////////////////

type ScanArgs struct {
	Cmd string
	Operand string
}

///////////////////////
// Application
///////////////////////

func init() {
	appChan = make(chan int)
}

type Application struct {
	backend Backend
}

// NewApp 创建Application对象。
func NewApp() Application {
	return Application {
		Backend {},
	}
}

// Run 启动并运行。
func (a Application) Run() int {
	fmt.Println("Running App ...")
	arg := ScanArgs {}
	mainLoop: for {
		fmt.Print("> ")
		fmt.Scanln(&arg.Cmd, &arg.Operand)
		switch arg.Cmd {
		case "eval":
			v, err := strconv.Atoi(arg.Operand)
			if err != nil {
				fmt.Printf("Invalid operand %q\n", arg.Operand)
			} else {
				go a.backend.heavyWork(v)
				fmt.Printf("Result:%d\n",<-appChan)
			}
			break
		case "exit":
			break mainLoop
		}
	}

	return 0
}

///////////////////////
// Backend
///////////////////////

type Backend struct {
	operate int
}

func NewBackend() Backend {
	return Backend {
		operate: EVAL,
	}
}

func (b Backend) run() {
	for b.operate != EXIT {
		switch b.operate {
		case EVAL:
			break
		case FETCH:
			break
		}
	}
}

func (b Backend) heavyWork(val int) {
	appChan <- (val * val)
}
