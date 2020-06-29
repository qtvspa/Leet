package main

import "Leet/src/base"

// 找出数组中第k大的元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 思路：1、递归分治 2、大根堆


func findKthLargest(nums []int, k int) int {
	heapSize := len(nums)

	// 构建大根堆
	base.BuildMaxHeap(nums, heapSize)

	// 进行 k-1次删除栈顶元素操作
	for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		base.MaxHeapify(nums, 0, heapSize)
	}
	// 此时栈顶元素即为目标元素
	return nums[0]
}

