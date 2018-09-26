package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/* 2-4 2-5 语句*/

func ifStatement() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)

	if err != nil { // 省略起始条件, for与if后面的条件没有括号
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

/*switch语句不用写break*/
func grade(score int) string {

	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score)) /*中断程序执行*/
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		g = "F"
	}
	return g
}

/*整数转二进制*/
func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/*输出文件中的每一行*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printContents(file)
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // 相当于while
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for { // 死循环
		fmt.Println("loop")
	}
}

func main() {

	ifStatement()
	fmt.Println(grade(0), grade(59), grade(99))

	// output: 101 1101
	fmt.Println(convertToBinary(5), convertToBinary(13))
	printFile("basic/branch/abc.txt")

	s := `ggg"k"gg
			tv	
			xq`
	printContents(strings.NewReader(s))
}
