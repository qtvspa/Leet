package main

import (
	"fmt"
	"strings"
)

// 验证一个字符是否是回文
// https://leetcode-cn.com/problems/valid-palindrome/
// 思路：双指针，同时要处理非字母数字的特殊字符


func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	left, right := 0, len(s) - 1

	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	// 验证字符串是否是字母or数字
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func main()  {
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Println(res)
}
