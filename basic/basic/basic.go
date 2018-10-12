package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*2-1 2-2 2-3 变量定义 常量 枚举 */

/*作用域是package包内部*/
var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a, b int = 3, 4
	var s string = "abc"

	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "str"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	/*第一次初始化用冒号*/
	a, b, c, s := 3, 4, true, "str"
	fmt.Println(a, b, c, s)
}

/*欧拉公式*/
func euler() {
	c := 3 + 4i
	fmt.Printf("euler: %f \n", cmplx.Abs(c))

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var a, b int = 5, 12
	fmt.Println(calcTriangle(a, b))
}
func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) /*强制类型转换*/
	return c
}

func consts() {
	//定义在包内、函数外也行，这样所有函数都能用
	const (
		filename string = "ggg.txt"
		a, b            = 3, 4
	)

	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

/*go语言的枚举*/
func enums() {
	// output : 0 1 2 3
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	// output : 0 2 3 4
	// iota可以作为自增值的种子
	const (
		php = iota
		_
		js
		ruby
		vb
	)

	// output: 1 1024 1048576 ....  (左移运算符)
	const (
		b  = 1 << (10 * iota) // 1 * 2^0
		kb                    // 1 * 2^10
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(php, js, ruby, vb)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	println("hello go")

	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	triangle()
	consts()
	enums()

	fmt.Println(aa, bb, ss)
}
