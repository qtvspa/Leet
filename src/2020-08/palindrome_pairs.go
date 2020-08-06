package main

// 字符列表寻找成组的回文
// https://leetcode-cn.com/problems/palindrome-pairs/
// 思路：哈希表


func palindromePairs(words []string) [][]int {

	revWords := [] string {}
	indices := map[string]int {}
	n := len(words)

	for _, word := range words{
		revWords = append(revWords, reverse(word))
	}

	for i := 0; i < n; i++ {
		indices[revWords[i]] = i
	}

	var ans [][]int

	for i := 0; i < n; i++ {
		if len(words[i]) == 0 {
			continue
		}
		for j := 0; j <= len(words[i]); j++ {
			if isPalindrome(words[i], j, len(words[i]) - 1){
				leftId := findWord(words[i], 0, j - 1, indices)
				if leftId != -1 && leftId != i {
					ans = append(ans, []int{i, leftId})
				}
			}
			if j != 0 && isPalindrome(words[i], 0, j - 1) {
				rightId := findWord(words[i], j,  len(words[i]) - 1, indices)
				if rightId != -1 && rightId != i {
					ans = append(ans, []int{rightId, i})
				}
			}
		}
	}

	return ans


}

func findWord(s string, left, right int, indices map[string]int) int {
	if v, ok := indices[s[left:right+1]]; ok {
		return v
	}
	return -1
}

func isPalindrome(s string, left, right int) bool {
	for i := 0; i < (right - left + 1) / 2; i++ {
		if s[left + i] != s[right - i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}