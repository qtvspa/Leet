package main

// 07-04 最长有效括号
// https://leetcode-cn.com/problems/longest-valid-parentheses/
// 思路：栈 or 动态规划


func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxAns = max(maxAns, i - stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
