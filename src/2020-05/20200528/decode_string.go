package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串解码
// https://leetcode-cn.com/problems/decode-string/
// 思路 栈


func decodeString(s string) string {

	var stk []string
	ptr := 0

	for ptr < len(s) {
		cur := s[ptr]

		if cur >= '0' && cur <= '9' {
			// 若当前字符为数字 则直接入栈
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			// 若当前字符为字母或[ 也直接入栈
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			var sub []string

			// 当遇到]时 执行出栈 直到遇到[时停止
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}

			// 将已出栈的字符反转
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// 当前栈顶为数字字符
			stk = stk[:len(stk)-1]

			// 字符转为数字 得到重复次数
			repTime, _ := strconv.Atoi(stk[len(stk)-1])

			// 数字出栈
			stk = stk[:len(stk)-1]

			// 按重复次数拼接字符串
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
		fmt.Println(stk)
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	// 提取字符串中的数据 并寻找到数字最后一位的下标

	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	// 将切片中的字符连接成字符串

	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}


func main() {
	a := decodeString("321[a]2[bc]")
	fmt.Println(a)
}

