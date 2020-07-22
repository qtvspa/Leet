package main

// 旋转数组的最小数
// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 思路：二分



func minArray(numbers []int) int {

	low := 0
	high := len(numbers) - 1

	for low < high {
		pivot := (low + high) / 2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}