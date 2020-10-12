package main

import (
	"math"
)

// 二叉排序树节点值的最小绝对值差
// https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
// 思路: 中序遍历

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt16, -1
	var search func(*TreeNode)

	search = func(node *TreeNode){
		if node == nil{
			return
		}

		search(node.Left)
		if pre > -1 && node.Val - pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		search(node.Right)
	}

	search(root)
	return ans
}
