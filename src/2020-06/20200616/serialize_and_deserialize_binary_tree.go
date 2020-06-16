package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 序列化和反序列化二叉树
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// 思路：深度优先 递归


// Definition for a binary tree node.
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

type Codec struct {
	l [] string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	// 整数转字符串 拼接
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}


func dfs(valsPtr *[]string) *TreeNode {
	// dfs

	val := (*valsPtr)[0]
	*valsPtr = (*valsPtr)[1:]

	if val == "#" {
		return nil
	}

	valInt, _ := strconv.Atoi(val)

	node := &TreeNode{Val: valInt}
	node.Left = dfs(valsPtr)
	node.Right = dfs(valsPtr)

	return node
}

func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	return dfs(&vals)
}

func main()  {
	 root := TreeNode{
	 	1,
	 	&TreeNode{
	 		2,
	 		nil,
	 		nil,
	 	},
	 	&TreeNode{
	 		3,
	 		&TreeNode{4, nil, nil},
	 		&TreeNode{5, nil, nil},
	 	},
	 }
	 obj := Constructor()

	 // 序列化
	 data := obj.serialize(&root)
	 fmt.Println(data)

	 // 反序列化
	 ans := obj.deserialize(data)
	 fmt.Println(ans)

}
