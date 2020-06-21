package main

import (
	"fmt"
	"math"
)

// 二叉树最大路径和
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 思路：使用递归求出每个节点的贡献值 同时更新最大路径和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int

	maxGain = func (node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右节点的最大贡献值
		leftMaxGain := max(maxGain(node.Left), 0)
		rightMaxGain := max(maxGain(node.Right), 0)

		//当前节点的最大路径和
		nodePathSum := node.Val + leftMaxGain + rightMaxGain

		// 比较并更新最大值
		maxSum = max(maxSum, nodePathSum)

		// 计算并返回当前节点的最大贡献值
		return max(leftMaxGain, rightMaxGain) + node.Val
	}

	maxGain(root)
	return maxSum
}


func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main()  {
	// testcase [-10,9,20,null,null,15,7]
	root := TreeNode{-10,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}

	res := maxPathSum(&root)
	fmt.Println(res)
}