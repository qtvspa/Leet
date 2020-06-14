package main

import "fmt"

// 顺时针打印矩阵
// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
// 思路 设定四个边界


func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var (
		rows, columns = len(matrix), len(matrix[0])
		order = make([]int, rows * columns)
		index = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)

	for left <= right && top <= bottom {
		// 打印从左到右的行
		for column := left; column <= right; column++ {
			order[index] = matrix[top][column]
			index++
		}
		// 打印右侧从上到下的列
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom {
			// 逆序打印下方的行
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			// 逆袭打印左侧的列
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}


func main() {
	a := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(spiralOrder(a))
}