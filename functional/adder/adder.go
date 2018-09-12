package main

import "fmt"

/*闭包 写法1*/
func adder() func(int) int {
	sum := 0
	println("param, sum", sum) // 这句根本不会执行，sum也不会每次进来变成0

	return func(v int) int {
		sum += v
		return sum
	}
}

/*闭包 写法2*/
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
