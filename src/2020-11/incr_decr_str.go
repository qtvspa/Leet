package main

import "fmt"

// 升降序重新排列字符串
// https://leetcode-cn.com/problems/increasing-decreasing-string/
// 思路：桶排序



func sortString(s string) string {
	// 初始化桶
	tmp := ['z' + 1]int{}
	for _, subS := range s{
		tmp[subS]++
	}

	n := len(s)
	answer := make([]byte, 0, n)

	for len(answer) < n {
		// 首先提取最长正序子串
		for i := byte('a'); i<= 'z'; i++ {
			if tmp[i] > 0 {
				answer = append(answer, i)
				tmp[i]--
			}
		}
		// 再提取最长倒序子串
		for i := byte('z'); i >= 'a'; i-- {
			if tmp[i] > 0 {
				answer = append(answer, i)
				tmp[i]--
			}
		}
	}

	return string(answer)
}

func main()  {
	ans := sortString("adaskjbrenfmzxmn")
	fmt.Println(ans)
}