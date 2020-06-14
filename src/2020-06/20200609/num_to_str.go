package main

import (
	"strconv"
)

// 把数字翻译成字符串
// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/

// 动态规划


func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1

	for i := 0; i < len(src); i++ {

		p, q, r = q, r, 0

		r += q

		if i == 0 {
			continue
		}

		pre := src[i-1:i+1]

		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}


func main()  {
	num := 12345
	translateNum(num)
}
