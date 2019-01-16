package main

import (
	"algorithm-demo/util"
	"fmt"
	"math"
)

/*
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
*/
func main() {
	fmt.Println(countNodes_2(util.ArrToTreeNode([]interface{}{1, 2, 3, 4, 5, 6, nil})))
	fmt.Println(countNodes_2(util.ArrToTreeNode([]interface{}{1, 2, 3, 4, nil, nil, nil})))
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

/////////////////////// Solution 1 ///////////////////////

func countNodes(root *util.TreeNode) int {
	step++
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

/////////////////////// Solution 2 ///////////////////////

/*
如果从某节点一直向左的高度 = 一直向右的高度, 那么以该节点为root的子树一定是complete binary tree.
而 complete binary tree的节点数,可以用公式算出 2^h - 1. 如果高度不相等, 则递归调用 return countNode(left) + countNode(right) + 1.  复杂度为O(h^2)
*/
func countNodes_2(root *util.TreeNode) int {
	step++
	if root == nil {
		return 0
	}

	leftHeight := leftHeight(root)
	rightHeight := rightHeight(root)

	if leftHeight == rightHeight {
		//return 1 + (2*leftHeight - 1) + countNodes(root.Left)
		return int(math.Pow(float64(2), float64(leftHeight))) - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
func leftHeight(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + leftHeight(root.Left)
}
func rightHeight(root *util.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + rightHeight(root.Right)
}
