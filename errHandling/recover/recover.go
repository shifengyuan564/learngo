package main

import (
	"fmt"
)

func tryRecover() {
	// 匿名函数
	defer func() {
		re := recover()
		if err, ok := re.(error); ok {
			fmt.Println("Error Occurred:", err)
		} else {
			panic(re)
		}
	}()

	b := 0
	a := 1 / b
	fmt.Println(a)

	//panic(errors.New("this is a custom error"))
}

func main() {
	tryRecover()
}
