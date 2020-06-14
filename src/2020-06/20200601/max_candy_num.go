package main

// 拥有最多糖果的孩子
// https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/


func kidsWithCandies(candies []int, extraCandies int) []bool {

	n := len(candies)
	maxCandyNum := 0

	for i := 0; i < n; i++ {
		maxCandyNum = max(candies[i], maxCandyNum)
	}
	result := make([]bool, n)

	for i := 0; i < n; i++ {
		result[i] = candies[i] + extraCandies >= maxCandyNum
	}

	return result

}


func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}
