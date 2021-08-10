package main

import (
	"fmt"
)

func main() {
	// 凑钱
	money := 15
	s := []int{1, 5, 11}
	fmt.Println("结果数组为", getLestNumForMoney(money, s))

	// 最长上升子序列
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

// 获取最长上升子序列
// 思路：f(x)=f(x-1)+x>x-1?x:忽略
func getLIS(s []int) (res []int) {
	res = []int{}

	return res
}
