package main

import "fmt"

// 找出给定字符串的最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 动态规划


func longestPalindrome(s string) string {
	n := len(s)
	ans := ""
	// 二维切片
	dp := make([][]int, n)
	// 将元素初始化为0值
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	fmt.Println(dp)

	for l := 0; l < n; l++ {
		for i := 0; i + l < n; i++ {
			j := i + l

			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] > 0 && l + 1 > len(ans) {
				ans = s[i:i+l+1]
			}
		}
	}
	fmt.Println(dp)
	return ans
}


func main()  {
	s := "asdfgfda"
	println(longestPalindrome(s))
}
