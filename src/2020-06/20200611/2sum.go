package main

import "fmt"

// 两数之和
// https://leetcode-cn.com/problems/two-sum/
// 思路：哈希表


func twoSum2(nums []int, target int) []int {

	m := map[int]int{}

	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}
	return nil
}

func main()  {
	a := []int{2,7,9,10}
	res := twoSum2(a, 9)
	fmt.Println(res)
}
