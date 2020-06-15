package main

import "fmt"

// 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/
// 思路：动态规划 使用滚动数组


func climbStairs(n int) int {
	p, q, result := 0, 0, 1

	for i:= 0; i < n ;i++ {
		p = q
		q = result
		result = p + q
	}
	return result
}

func main()  {
	fmt.Println(climbStairs(5))
}