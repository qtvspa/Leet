package main

import "fmt"

// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 思路 快慢指针

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] { }
	fmt.Println(slow, fast)
	slow = 0

	for slow != fast {
		fmt.Println(slow, fast)
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main ()  {
	a := []int {3,2,1,2,3}
	fmt.Println(findDuplicate(a))
}
