package main


// 判断输入的整形数是否为回文数
// https://leetcode-cn.com/problems/palindrome-number/


func isPalindrome(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	// 如何知道反转数字的位数已经达到原始数字位数的一半？
	// 由于整个过程我们不断将原始数字除以 10，然后给反转后的数字乘上 10，
	// 所以，当原始数字小于或等于反转后的数字时，就意味着我们已经处理了一半位数的数字了。
	revertedNumber := 0
	for x > revertedNumber {
		// x每次除10求余数的目的是 获取个十百千万..每位的单个数字
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber / 10
}
