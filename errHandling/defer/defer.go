package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println("defer", 1)
	fmt.Println("defer", 2)

	for i := 0; i < 50; i++ {
		defer fmt.Println(i)

		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return

		//panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 从内存写到文件

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, i)
	}
}

func main() {
	tryDefer()

	writeFile("GoWriteFile.txt")
}
