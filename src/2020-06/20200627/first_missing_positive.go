package main

// 数组中第一个缺失的正整数
// https://leetcode-cn.com/problems/first-missing-positive/
// 思路：用数组的索引做哈希表的key
// 1、遍历+做标记 2、遍历+置换位置


// 1、做标记
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 首先把所有非正整数置为 n+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 给所有 <= n 的数 其值对应的索引位置上的数打上标记 即：变负数
	for i := 0; i < n; i++ {
		num := abs(nums[i])

		if num <= n {
			nums[num - 1] = -abs(nums[num - 1])
		}
	}

	// 返回第一个正数所在的索引位置 + 1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	// 如果数组元素均为负数 则返回 n + 1
	return n + 1
}


// 2、置换
func firstMissingPositive2(nums []int) int {
	n := len(nums)
	// 对数组进行一次遍历 对于遍历到的数 x=nums[i]
	// 如果 x∈[1,n] 就知道x应当出现在数组中的 x−1 的位置
	// 此时交换 nums[i] 和 nums[x−1] 这样x就出现在了正确的位置
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	// 返回第一个位置错误的数
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	// 如果数组元素位置均正确 返回 n+1
	return n + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}