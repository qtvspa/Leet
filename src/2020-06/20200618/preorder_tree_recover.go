package main

import "fmt"

// 从先序遍历还原二叉树
// https://leetcode-cn.com/problems/recover-a-tree-from-preorder-traversal/
// 思路：首先提取每个节点的深度信息和值信息，然后根据已有的节点数和当前节点的深度来判定位置（左还是右）


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {

	var path []*TreeNode

	pos := 0

	for pos < len(S) {

		// 提取深度信息
		level := 0
		for S[pos] == '-' {
			level++
			pos++
		}

		// 提取值信息
		value := 0
		for ; pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos++ {
			value = value * 10 + int(S[pos] - '0')
		}

		node := &TreeNode{Val: value}

		if level == len(path) {
			// 当前节点为上一节点的左节点
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			// 当前节点为对应深度的那个父节点的右节点
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}

	return path[0]
}

// 先序遍历
func PreTrans(root * TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreTrans(root.Left)
	PreTrans(root.Right)
}

func main(){
	s := "1-2--3--4-5--6--7"
	res := recoverFromPreorder(s)
	PreTrans(res)
}
