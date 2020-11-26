package main

import (
	"fmt"
)

// 求数组元素的最大间距
// https://leetcode-cn.com/problems/maximum-gap/
// 思路：桶排序

type pair struct{ min, max int }

func maximumGap(nums []int) (ans int) {
	n := len(nums)
	if n < 2 {
		return
	}

	minVal := min(nums...)
	maxVal := max(nums...)

	// 计算桶的个数 以及每个桶的长度
	d := max(1, (maxVal-minVal)/(n-1))
	bucketSize := (maxVal-minVal)/d + 1

	// 初始化桶，其中存储 (桶内最小值，桶内最大值) 对，(-1, -1) 表示该桶是空的
	buckets := make([]pair, bucketSize)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	// 将数组中元素添加至桶中 同时计算每个桶的桶内最大值和最小值
	for _, v := range nums {
		bid := (v - minVal) / d
		if buckets[bid].min == -1 {
			buckets[bid].min = v
			buckets[bid].max = v
		} else {
			buckets[bid].min = min(buckets[bid].min, v)
			buckets[bid].max = max(buckets[bid].max, v)
		}
	}

	// 计算桶与桶之间的最大差值
	prev := -1
	for i, b := range buckets {
		if b.min == -1 {
			continue
		}
		if prev != -1 {
			ans = max(ans, b.min-buckets[prev].max)
		}
		prev = i
	}
	return
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func main()  {
	a := []int{2,1,3,4,6,9,11,5}
	result := maximumGap(a)
	fmt.Println(result)
}