package main

import (
	"algorithm-demo/util"
	"fmt"
)

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
*/
func main() {
	sum := hasPathSum(util.ArrToTreeNode([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, nil, nil, 1}), 22)
	println(sum)
	fmt.Printf("step=%d", step)
}

var step int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *util.TreeNode, sum int) bool {
	step++
	if root == nil {
		return false
	}
	//TODO 一定要注意：必须要到叶子节点
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
