package main

import (
	"algorithm-demo/util"
	"log"
	"reflect"
)

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3,4]
   1
    \
     2
    / \
   3   4

输出: [1,2,3,4]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func main() {
	arr := preorderTraversak_3(&util.TreeNode{Val: 1, Left: nil, Right: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 3}, Right: &util.TreeNode{Val: 4}}})
	log.Println(arr)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//Solution1==> 递归方式，使用切片方式
//时间复杂度：O(n)
//空间复杂度：O(h)
func preorderTraversal(root *util.TreeNode) []int {
	//先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)
	//中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
	//后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

	var res []int
	if root != nil {
		res = append(res, root.Val)
		leftArr := preorderTraversal(root.Left)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(leftArr)).Interface().([]int)
		rightArr := preorderTraversal(root.Right)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(rightArr)).Interface().([]int)
	}
	return res
}

//Solution1==> 迭代方式
//时间复杂度：O(n)
//空间复杂度：O(h)
func preorderTraversal_2(root *util.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := &util.Stack{}
	//1. 使用栈存放所有节点，初始压入根节点
	stack.Push(root)
	//2. 使用循环每次从栈顶取出一个节点，打印该节点的value，然后依次将右节点和做节点压入栈顶（循环条件是栈不为空）
	for !stack.Empty() {
		node := stack.Top().(*util.TreeNode)
		//打印节点值
		res = append(res, node.Val)
		stack.Pop()
		//压入右节点
		if node.Right != nil {
			stack.Push(node.Right)
		}
		//压入左节点
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}

	return res
}

//Solution1==> morris解法
//时间复杂度：O(n)
//空间复杂度：O(1)

//TODO （待完善） morris解法的算法描述如下：
//1、如果当前结点的左孩子为空，则输出当前结点并将当前结点的右结点作为当前结点。
//2、如果当前结点的左孩子不为空，则从当前结点的左子树找出当前结点的前去节点：
//如果前驱结点p的右孩子为空，则将p的右孩子设为当前结点；否则，输出当前结点，并将p的右孩子置为空，并将当前当前结点的右孩子置为当前结点
//3、重复1 ，2两步直到当前结点为空
func preorderTraversak_3(root *util.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			res = append(res, cur.Val)
			cur = cur.Right
		} else {
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			if prev.Right != nil {
				prev.Right = nil
				res = append(res, cur.Val)
				cur = cur.Right
			} else {
				prev.Right = cur
				cur = cur.Left
			}
		}
	}
	return res
}
