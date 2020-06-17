package main

import "fmt"

// 最佳观光组合（智力题）
// https://leetcode-cn.com/problems/best-sightseeing-pair/
// 思路 一次循环即可


func maxScoreSightseeingPair(A []int) int {

	sum_i := A[0]
	answer := 0

	for j := 1; j < len(A); j++ {
		answer = max(A[j] - j + sum_i, answer)
		sum_i = max(sum_i, A[j] + j)
	}

	return answer
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main()  {
	a := [] int {8, 1, 5, 2, 6}
	res := maxScoreSightseeingPair(a)
	fmt.Println(res)
}