package main

// 新21点
// https://leetcode-cn.com/problems/new-21-game/

//爱丽丝以 0 分开始，并在她的得分少于 K 分时抽取数字。
//抽取时，她从 [1, W] 的范围中随机获得一个整数作为分数进行累计，其中 W 是整数。 每次抽取都是独立的，其结果具有相同的概率。
//当爱丽丝获得不少于 K 分时，她就停止抽取数字。 爱丽丝的分数不超过 N 的概率是多少？



func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K + W)
	for i := K; i <= N && i < K + W; i++ {
		dp[i] = 1.0
	}
	for i := K - 1; i >= 0; i-- {
		for j := 1; j <= W; j++ {
			dp[i] += dp[i + j] / float64(W)
		}
	}
	return dp[0]
}

func new21Game2(N int, K int, W int) float64 {
	if K == 0 {
		return 1.0
	}
	dp := make([]float64, K + W)
	for i := K; i <= N && i < K + W; i++ {
		dp[i] = 1.0
	}

	dp[K - 1] = 1.0 * float64(min(N - K + 1, W)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i + 1] - (dp[i + W + 1] - dp[i + 1]) / float64(W)
	}
	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

