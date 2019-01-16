package main

import "algorithm-demo/util"

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/
func main() {
	treeNode := util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	println(minDepth(treeNode))
	treeNode2 := &util.TreeNode{Val: 1, Left: &util.TreeNode{Val: 2}}
	println(minDepth(treeNode2))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	//叶子节点是指没有子节点的节点
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + util.Min(minDepth(root.Left), minDepth(root.Right))
}
