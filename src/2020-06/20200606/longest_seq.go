package main

import "fmt"

// 给定一个未排序的整数数组，找出最长连续序列的长度。
// https://leetcode-cn.com/problems/longest-consecutive-sequence/


func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}

	for _, num := range nums{
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet{

		if !numSet[num-1]{

			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1]{
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak

}



func main()  {
	nums := []int{2,5,6,3,10,5}
	fmt.Println(longestConsecutive(nums))
}
