package base

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 前序遍历
func PreTrans(root * TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreTrans(root.Left)
	PreTrans(root.Right)
}

// 中序遍历
func MidTrans(root *TreeNode) {
	if root == nil {
		return
	}

	MidTrans(root.Left)
	fmt.Println(root)
	MidTrans(root.Right)
}

//后序遍历:先遍历左子树再遍历根，再遍历右子树，再遍历根节点
func PostTrans(root *TreeNode) {
	if root == nil {
		return
	}
	PostTrans(root.Left)
	PostTrans(root.Right)
	fmt.Println(root.Val)
}