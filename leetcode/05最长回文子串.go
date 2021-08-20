package main

import (
	"fmt"
)

func main() {
	// 凑钱
	//fmt.Println("结果数组为:", getLestNumForMoney(15, []int{1, 5, 11}))

	// 回文子串数量
	s1 := "abcccba"
	fmt.Println(fmt.Sprintf("%s的回文数量:%d", s1, getHuiWenNumber(s1)))
	fmt.Println(fmt.Sprintf("%s的回文数量(双指针):%d", s1, getHuiWenNumberV2(s1)))
}

// 典型的动态规划，尽量少的张数n凑出足额的钱w，给定面额数组s
// 思路：f(n)=min{f(n-s1),f(n-s2),f(n-s3),...}+1
func getLestNumForMoney(w int, s []int) (res []int) {
	res = []int{}
	res = append(res, 0)
	for i := 1; i <= w; i++ {
		cost := w
		for _, v := range s {
			if i-v >= 0 {
				cost = min(cost, res[i-v]+1)
			}
		}
		res = append(res, cost)
	}
	return res
}

// 获取一段字符串的回文数量
func getHuiWenNumber(str string) (res int) {
	s := []byte(str)
	dp := make([][]bool, len(s))
	// 不定长二维数组初始化
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	// 算法
	for i := len(s) - 1; i >= 0; i-- { // 注意遍历顺序
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}

	}
	fmt.Println(dp)
	return res
}

// 获取一段字符串的回文数量,双指针法
func getHuiWenNumberV2(str string) (res int) {
	s := []byte(str)
	for i := 0; i < len(s); i++ {
		// 为中心
		j1 := i
		j2 := i
		for j1 >= 0 && j2 < len(s) {
			if s[j1] == s[j2] {
				res++
			} else {
				break
			}
			j1--
			j2++
		}
		j1 = i
		j2 = i + 1
		for j1 >= 0 && j2 < len(s) {
			if s[j1] == s[j2] {
				res++
			} else {
				break
			}
			j1--
			j2++
		}

	}
	return res
}

// 比大小
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
