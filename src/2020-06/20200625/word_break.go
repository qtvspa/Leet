package main

import "fmt"

// 单词拆分
// https://leetcode-cn.com/problems/word-break/
// 思路：动态规划


func wordBreak(s string, wordDict []string) bool {

	wordDictSet := make(map[string]bool)

	// 初始化哈希表 用于后续的【判断一个字符串是否在此哈希表中】
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	// 初始化dp数组
	// db[i]表示 字符串s的前i个字符 s[0:i-1]能否被空格拆分成若干个字典中的单词
	dp := make([]bool, len(s) + 1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			fmt.Println(dp, i, j)
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main(){
	ss := "leetcode"
	words := []string{"leet", "code"}
	res := wordBreak(ss, words)
	fmt.Println(res)
}