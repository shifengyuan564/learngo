package main

import "fmt"

/*遍历map, 内建类型都可以作为key */
func traversingMap() {
	m := map[string]string{
		"name":   "jackson",
		"course": "golang",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	/*通过key获取value*/
	courseName, ok := m["course"]
	if ok {
		fmt.Println("课程名称：", courseName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	delete(m, "name")
	fmt.Println("After delete key 'name': ", m)

}

func main() {
	/*map中的值是无序的, hashmap*/
	m := map[string]string{
		"name":   "jackson",
		"course": "golang",
	}

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      // m3 == nil ， empty map == nil

	fmt.Println(m, m2, m3)

	traversingMap()
}
