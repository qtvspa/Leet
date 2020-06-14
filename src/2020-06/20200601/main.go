package main

import "fmt"

func main()  {


	//candies := []int {1,2,2,3,1}
	//fmt.Println(kidsWithCandies(candies, 1))

	Node15 := TreeNode{Val: 15}
	Node14 := TreeNode{Val: 14}
	Node13 := TreeNode{Val: 13}
	Node12 := TreeNode{Val: 12}
	Node11 := TreeNode{Val: 11}
	Node10 := TreeNode{Val: 10}
	Node9 := TreeNode{Val: 9}
	Node8 := TreeNode{Val: 8}
	Node7 := TreeNode{Val: 7, Left:&Node14, Right:&Node15}
	Node6 := TreeNode{Val: 6, Left:&Node12, Right:&Node13}
	Node5 := TreeNode{Val: 5, Left:&Node10, Right:&Node11}
	Node4 := TreeNode{Val: 4, Left:&Node8, Right:&Node9}
	Node3 := TreeNode{Val: 3, Left:&Node6, Right:&Node7}
	Node2 := TreeNode{Val: 2, Left:&Node4, Right:&Node5}
	Node1 := TreeNode{Val: 1, Left:&Node2, Right:&Node3}

	resultNode := lowestCommonAncestor(&Node1, &Node11, &Node9)
	fmt.Println(resultNode.Val)


}
