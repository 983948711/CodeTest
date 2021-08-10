package main

import "fmt"

func main() {
	// data还是array
	data := [...]int{0, 1, 2, 3, 4, 5}

	// s就是slice
	s := data[2:4]
	s = append(s, 3)
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)

	// 创建slice
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}
