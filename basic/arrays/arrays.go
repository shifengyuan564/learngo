package main

import "fmt"

/* 3-1 数组 */

func useRange() {

	var arr1 [5]int
	arr2 := [3]int{7, 8, 9}
	arr3 := [...]int{2, 4, 6, 8, 10}

	/*二维数组，4行5列*/
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	/*遍历 range关键字可以获得下标、值*/
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}
}

/* []int是切片，引用类型；[5]int是数组，值类型 */
func printArray(arr [5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/*指针传值*/
func printArrayPtr(arr *[5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//  数组是值类型，赋值和函数传参操作都会复制整个数组数据。而原数组不会变化
func arrayCopy() {
	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA
	arrayB[1] = 300

	fmt.Printf("ArrayCopy Test : arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("ArrayCopy Test : arrayB : %p , %v\n", &arrayB, arrayB)

	testArrayCopy(arrayA)
}

func testArrayCopy(x [2]int) {
	x[1] = 400
	fmt.Printf("ArrayCopy Test : func Array : %p , %v\n", &x, x)
}

func main() {

	//useRange()

	arr5 := [...]int{2, 4, 6, 8, 10}
	printArray(arr5) // output: 100,4,6,8,10

	for i, v := range arr5 {
		fmt.Println(i, v) // output: 2,4,6,8,10 （证明了数组参数是值传递，无法改变原值）
	}

	// Test Case 1:
	printArrayPtr(&arr5) // 指针传递可以改变原数组的值

	// Test Case 2: 打印出的三个内存地址都不同，验证了Go中数组赋值和函数传参都是值复制
	arrayCopy()
}
