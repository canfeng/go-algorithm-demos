package main

import "algorithm-demo/util"

/*
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
func main() {
	treeNode := util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	println(maxDepth_2(treeNode))
	//treeNode1 := util.ArrToTreeNode([]interface{}{})
	//println(maxDepth(treeNode1))
	//treeNode2 := util.ArrToTreeNode([]interface{}{1, 2})
	//println(maxDepth(treeNode2))
}

type NodeItem struct {
	depth int
	node  *util.TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	return curMaxDepth(&NodeItem{node: root, depth: 1})
}

func curMaxDepth(nodeItem *NodeItem) int {
	node := nodeItem.node
	depth := nodeItem.depth
	if node.Left == nil && node.Right == nil {
		return depth
	}
	var leftDepth, rightDepth int
	if node.Left != nil {
		leftDepth = curMaxDepth(&NodeItem{node: node.Left, depth: depth + 1})
	}
	if node.Right != nil {
		rightDepth = curMaxDepth(&NodeItem{node: node.Right, depth: depth + 1})
	}
	return util.Max(leftDepth, rightDepth)
}

//////////////////////// Solution 2 /////////////////////////

func maxDepth_2(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + util.Max(maxDepth_2(root.Left), maxDepth_2(root.Right))
}
