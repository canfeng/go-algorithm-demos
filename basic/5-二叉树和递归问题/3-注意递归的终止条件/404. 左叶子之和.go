package main

import (
	"algorithm-demo/util"
	"fmt"
)

/*
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/
func main() {
	res := sumOfLeftLeaves_2(util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}))
	println(res)
	fmt.Printf("s1=%d\n", s1)
	res = sumOfLeftLeaves_2(util.ArrToTreeNode([]interface{}{1, 2, 3, 4, 5, nil, nil}))
	println(res)

	fmt.Printf("s1=%d\n", s1)
}

var res int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

////////////////////////// Solution 1  //////////////////////////

func sumOfLeftLeaves(root *util.TreeNode) int {
	res = 0
	calcSum(root)
	return res
}

func calcSum(root *util.TreeNode) {
	if root == nil {
		return
	}
	//TODO 注意是 左叶子节点
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	calcSum(root.Left)
	calcSum(root.Right)
}

////////////////////////// /Solution 2 //////////////////////////////
var s1 int

func sumOfLeftLeaves_2(root *util.TreeNode) int {

	s1++

	ret := 0

	if root == nil {
		return 0
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			ret += root.Left.Val
		}
		ret += sumOfLeftLeaves_2(root.Left)
	}
	ret += sumOfLeftLeaves_2(root.Right)

	return ret
}
