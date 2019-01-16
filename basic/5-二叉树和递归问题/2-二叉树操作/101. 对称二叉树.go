package main

import (
	"algorithm-demo/util"
	"container/list"
)

/*
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
func main() {
	println(isSymmetric(util.ArrToTreeNode([]interface{}{1, 2, 2, 3, 4, 4, 3})))
	println(isSymmetric(util.ArrToTreeNode([]interface{}{1, 2, 2, nil, 3, nil, 3})))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *util.TreeNode) bool {
	if root == nil {
		return true
	}
	return iterationSolution(root.Left, root.Right)
}

/*
如果同时满足下面的条件，两个树互为镜像：

1. 它们的两个根结点具有相同的值。
2. 每个树的右子树都与另一个树的左子树镜像对称。
*/

/*
递归解法
*/
func recursionSolution(node1 *util.TreeNode, node2 *util.TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil || node1.Val != node2.Val {
		return false
	}

	return recursionSolution(node1.Left, node2.Right) && recursionSolution(node1.Right, node2.Left)
}

/*
迭代解法
*/
func iterationSolution(node1 *util.TreeNode, node2 *util.TreeNode) bool {
	//使用队列
	lq := list.New()
	rq := list.New()
	lq.PushBack(node1)
	rq.PushBack(node2)
	for lq.Len() > 0 && rq.Len() > 0 {
		front := lq.Front()
		node1 = front.Value.(*util.TreeNode)
		lq.Remove(front)

		front = rq.Front()
		node2 = front.Value.(*util.TreeNode)
		rq.Remove(front)

		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		lq.PushBack(node1.Left)
		lq.PushBack(node1.Right)

		rq.PushBack(node2.Right)
		rq.PushBack(node2.Left)
	}
	return true
}
