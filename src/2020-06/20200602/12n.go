package main

// 求1+2+...+n的和
// https://leetcode-cn.com/problems/qiu-12n-lcof/


func sumNums(n int) int {
	result := 0
	
	var sum func(int) bool
	sum = func(n int) bool {
		result += n
		return n > 0 && sum(n - 1)
	}

	sum(n)

	return result
}