package main

import (
	"fmt"
	"sort"
)

// 按要求转变数组 使其转变后最接近目标值的数组和
// https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/
// 思路：二分查找+枚举


func findBestValue(arr []int, target int) int {

	// 排序
	sort.Ints(arr)

	n := len(arr)

	// 计算累加和 并保存
	prefix := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	// 求数组最大值
	r := arr[n-1]

	ans, diff := 0, target

	// 进行枚举
	// 同时计算每个枚举条件下的数组和与target的差值（取绝对值）
	for i := 1; i <= r; i++ {

		// 寻找当前枚举值在已排序数组中的索引位置
		index := sort.SearchInts(arr, i)

		if index < 0 {
			index = -index - 1
		}
		// 计算当前枚举条件下的数组和（根据题意，大于当前枚举值的数需要被替换成枚举值）
		cur := prefix[index] + (n - index) * i

		// 比较diff 选出diff最小时所对应的那个数
		if abs(cur - target) < diff {
			ans = i
			diff = abs(cur - target)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main()  {
	a := [] int {1,5,3,4}
	fmt.Println(findBestValue(a, 12))
}