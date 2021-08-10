package main

import "fmt"

func main() {
	// 1 两数之和
	nums := []int{2, 11, 7, 15}
	target := 9
	fmt.Println("1求和", getSumNumber(nums, target))
	// 这里主要考察哈希表的查找
	// 242.有效的字母异位符
	s1 := "helalo"
	s2 := "elhloa"
	fmt.Println("242异位数结果是,方法1：", isAnagram1(s1, s2))
	fmt.Println("242异位数结果是,方法2：", isAnagram2(s1, s2))
	// 394.两个数组的交集
	num1 := []int{4, 9, 5, 6, 6, 3}
	num2 := []int{2, 2, 3, 4, 4, 6, 6, 6, 20}
	fmt.Println("394两个数组的交集结果是", getIntersection(num1, num2))
}

// 获取数组两数之和为某个数的数组下表,只有唯一值
func getSumNumber(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}

// 判断两个数是否是异位数，用map查找
func isAnagram1(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	existsMap := make(map[byte]int)
	// 构造map
	for i := 0; i < len(s); i++ {
		if _, ok := existsMap[s[i]]; ok {
			existsMap[s[i]]++
		} else {
			existsMap[s[i]] = 1
		}
	}

	for i := 0; i < len(s); i++ {
		if v, ok := existsMap[s[i]]; v >= 1 && ok {
			existsMap[s[i]]--
		} else {
			return false
		}
	}
	return true
}

// 这里用数组查找，因为26个英文字母是有限的
func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	exist := [26]int{}
	// 构造map
	for i := 0; i < len(s); i++ {
		exist[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		exist[t[i]-'a']--
	}
	for i := 0; i < len(exist); i++ {
		if exist[i] != 0 {
			return false
		}
	}
	return true
}

// 获取数组的交集
func getIntersection(s1, s2 []int) []int {
	var res []int
	if len(s1) == 0 || len(s2) == 0 {
		return res
	}
	existsMap := make(map[int]bool)
	for _, v := range s1 {
		existsMap[v] = false
	}

	for _, v := range s2 {
		if isOut, ok := existsMap[v]; ok && !isOut {
			res = append(res, v)
			existsMap[v] = true
		}
	}
	return res
}
