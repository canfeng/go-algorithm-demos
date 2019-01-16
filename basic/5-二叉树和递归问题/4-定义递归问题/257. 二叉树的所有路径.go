package main

import (
	"algorithm-demo/util"
	"fmt"
	"strconv"
)

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/
func main() {
	println(binaryTreePaths(util.ArrToTreeNode([]interface{}{1, 2, 3, nil, 5, nil, nil})))
}

func binaryTreePaths(root *util.TreeNode) []string {
	var res [][]int

	cur(root, res)

	return res
}

func cur(root *util.TreeNode, res [][]int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)

	cur(root.Left, res)
	cur(root.Right, res)
}
