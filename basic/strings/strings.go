package main

import "fmt"

func main() {
	s := "Yes!学习区块链"              // UTF-8
	fmt.Println("字符串长度：", len(s)) // 19

	// 59 65 73 21 E5 AD A6 E4 B9 A0 E5 8C BA E5 9D 97 E9 93 BE
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// (0 59) (1 65) (2 73) (3 21) (4 5B66) (7 4E60) (10 533A) (13 5757) (16 94FE)
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// (0 Y) (1 e) (2 s) (3 !) (4 学) (5 习) (6 区) (7 块) (8 链)
	/* 使用range遍历rune, rune相当于go的char */
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
