package main

import (
	"math"
	"sort"
)

// 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 思路：前缀和+二分


func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n + 1)

	// 计算前缀和数组
	for i := 1; i <= n; i++ {
		sums[i] = sums[i - 1] + nums[i - 1]
	}

	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound - (i - 1))
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
