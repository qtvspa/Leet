package main


// 正则表达式匹配
// https://leetcode-cn.com/problems/regular-expression-matching/
// 思路：动态规划



func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// 用于判断s和p对应位置的单个字符是否匹配
	matches := func(i, j int) bool {

		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m + 1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n + 1)
	}

	f[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果p的第j个字符是 *，那么就表示可以对p的第 j-1 个字符匹配任意自然数次
			// 此时需要判定的是f[i][j]和f[i][j-2]是否相等
			if p[j-1] == '*' {

				f[i][j] = f[i][j] || f[i][j-2]

				if matches(i, j - 1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}

			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
