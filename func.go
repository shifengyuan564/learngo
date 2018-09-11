package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/*2-6 2-7 函数 数组*/

/*a b 同类型*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := divi(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func divi(a, b int) (q, r int) {
	return a / b, a % b
}

/*op是函数类型的参数，含有两个int参数，返回值是int; a,b是另外两个参数*/
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d) \n", opName, a, b)

	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/*可变参数列表*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	/**b, *a = *a, *b*/
	b, a = a, b
	return a, b
}

func main() {
	fmt.Println(eval(17, 2, "/"))
	fmt.Println(eval(17, 2, "x"))

	q, r := divi(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	/*指针*/
	a, b := 9, 10
	a, b = swap(a, b)
	fmt.Println(a, b)
}
