package main

import "fmt"

var (
	s         []byte      // 存放单个结果
	res       []string    // 存放所有结果
	letterMap = []string{ // 映射
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
)

func main() {
	params := []int{2, 3, 4}
	backTracking(params, 0)
	fmt.Println(res)
}

// params表示给定的号码数组，index表示遍历的第几个数字
func backTracking(params []int, index int) {
	// 终止条件
	if index == len(params) {
		res = append(res, string(s))
		return
	}
	// 单层遍历
	number := params[index]              // 当前号码
	letters := []byte(letterMap[number]) // 号码对应的字符串
	for i := 0; i < len(letters); i++ {
		s = append(s, letters[i])
		backTracking(params, index+1)
		s = s[0 : len(s)-1]
	}
}

/** 回溯法解决，回溯三部曲
1、递归函数的参数以及返回值
2、回溯函数终止条件
3、单层遍历
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
**/
