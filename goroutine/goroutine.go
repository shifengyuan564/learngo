package main

import (
	"fmt"
	"time"
)

func main() {

	// 开了1000个goroutine协程。其他的语言通过线程是无法做到的，需要用异步IO才能达到这个效果
	for i := 0; i < 1000; i++ {
		go func(id int) {
			for {
				fmt.Printf("Hello from goroutine %d\n", id)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
