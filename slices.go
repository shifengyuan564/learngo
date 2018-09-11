package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

/*输出slice的长度和容量*/
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func createSlice() {
	var s []int // 未赋值时候s == nil
	printSlice(s)

	for i := 0; i < 100; i++ {
		s = append(s, i)
		printSlice(s)
	}

	fmt.Println(s)
}

func copySlice() {
	s1 := []int{9, 8, 7, 6, 5}
	s2 := make([]int, 10, 32)
	copy(s2, s1)
	printSlice(s2)
}

func deleteElementFromSlice() {
	s1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	s1 = append(s1[:3], s1[4:]...) // 删除下标为3的元素
	printSlice(s1)                 // output: [9 8 7 5 4 3 2 1]
}

func poping() {
	s1 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	// From Front
	front := s1[0]
	s1 = s1[1:]

	// From tail
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]

	fmt.Println("Poping Front & tail :", front, tail) // 9 1
	printSlice(s1)                                    // output: len=7, cap=8, [8 7 6 5 4 3 2]
}

func appendSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	sli := arr[:] // 转换成slice
	arr1 := append(sli, 8)
	arr2 := append(sli, 9)

	fmt.Println("arr,sli,arr1,arr2 : ", arr, sli, arr1, arr2)
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7} // 此时是一个数组，而不是切片。[]int{1,2,3,4}是切片
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	fmt.Println("\n After updateSlice(s1) ")
	s1 := arr[2:]
	updateSlice(s1)  /* 会将arr，s1都给改掉。用slice就不用指针了 */
	fmt.Println(s1)  // [100 3 4 5 6 7]
	fmt.Println(arr) // [0 1 100 3 4 5 6 7]

	s1 = s1[4:]     /* reslice */
	fmt.Println(s1) // [6 7]

	appendSlice()
	createSlice()
	copySlice()
	deleteElementFromSlice()
	poping()
}
