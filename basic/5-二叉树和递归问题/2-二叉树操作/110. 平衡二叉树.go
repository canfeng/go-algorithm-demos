package main

import (
	"algorithm-demo/util"
	"fmt"
	"math"
)

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func main() {
	step = 0
	println(isBalanced_2(util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})))
	fmt.Printf("step=%d\n", step)
	step = 0
	println(isBalanced_2(util.ArrToTreeNode([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4})))
	fmt.Printf("step=%d\n", step)
}

var step int

/////////////////////////// Solution 1 /////////////////////////////

//时间复杂度：O(nlogn)
//空间复杂度：O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *util.TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalancedNode(root)
}

func isBalancedNode(root *util.TreeNode) bool {
	if root == nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if math.Abs(float64(lh)-float64(rh)) <= 1 {
		return isBalancedNode(root.Left) && isBalancedNode(root.Right)
	}
	return false
}

func height(root *util.TreeNode) int {
	step++
	if root == nil {
		return 0
	}
	return 1 + util.Max(height(root.Left), height(root.Right))
}

/////////////////////////// Solution 2 /////////////////////////////
func isBalanced_2(root *util.TreeNode) bool {
	balanced, _ := calcHeight(root)
	return balanced
}

func calcHeight(root *util.TreeNode) (bool, int) {
	step++
	if root == nil {
		return true, 0
	}
	lb, lh := calcHeight(root.Left)
	rb, rh := calcHeight(root.Right)
	height := math.Abs(float64(lh) - float64(rh))
	if lb && rb && height <= 1 {
		return true, 1 + util.Max(lh, rh)
	}
	return false, -1
}
