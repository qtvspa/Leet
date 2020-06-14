package main

// https://leetcode-cn.com/problems/product-of-array-except-self/
// 给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，
// 其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。（禁用除法）

// 思路：计算出两个数组 一个计算从左往右的乘积 一个计算从右往左的乘积


func productExceptSelf(nums []int) []int {

	length := len(nums)
	L, R, ans := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i - 1] * L[i - 1]
	}

	R[length - 1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i + 1] * R[i + 1]
	}

	for i := 0; i < length; i++ {
		ans[i] = L[i] * R[i]
	}

	return ans
}