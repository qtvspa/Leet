package main

import "fmt"

// https://leetcode-cn.com/problems/valid-palindrome-ii/


// 工具函数 判断一个纯英文字符串是否为回文字符串
func isPalindrome(s string) bool {
	start, end := 0, len(s) - 1
	for start < end {
		if s[start] == s[end]{
			start++
			end --
		} else{
			return false
		}
	}
	return true
}


func validPalindrome(s string) bool {

	start, end := 0, len(s) - 1

	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return isPalindrome(s[start:end]) || isPalindrome(s[start+1:end+1])
		}
	}
	return true
}

func main()  {
	a := "abac"
	fmt.Println(validPalindrome(a))
}
