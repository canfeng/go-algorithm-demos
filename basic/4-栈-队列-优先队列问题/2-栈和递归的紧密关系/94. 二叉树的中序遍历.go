package main

import (
	"algorithm-demo/util"
	"log"
	"reflect"
)

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/
func main() {
	arr := inorderTraversal_2(&util.TreeNode{Val: 1, Left: nil, Right: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 3}, Right: &util.TreeNode{Val: 4}}})
	log.Println(arr)
}

//递归实现
func inorderTraversal_1(root *util.TreeNode) []int {
	//中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
	var res []int
	if root != nil {
		leftArr := inorderTraversal_1(root.Left)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(leftArr)).Interface().([]int)

		res = append(res, root.Val)

		rightArr := inorderTraversal_1(root.Right)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(rightArr)).Interface().([]int)
	}

	return res
}

func inorderTraversal_2(root *util.TreeNode) []int {
	//中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
	var res []int
	if root == nil {
		return res
	}
	stack := &util.Stack{}
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}

		node = stack.Top().(*util.TreeNode)
		stack.Pop()
		res = append(res, node.Val)
		node = node.Right
	}

	return res
}
