package main

import (
	"algorithm-demo/util"
	"log"
	"reflect"
)

/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func main() {
	arr := postorderTraversal_2(&util.TreeNode{Val: 1, Left: nil, Right: &util.TreeNode{Val: 2, Left: &util.TreeNode{Val: 3}, Right: &util.TreeNode{Val: 4}}})
	log.Println(arr)
}

//后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)

//递归方式
func postorderTraversal(root *util.TreeNode) []int {
	var res []int
	if root != nil {
		leftArr := postorderTraversal(root.Left)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(leftArr)).Interface().([]int)

		res = append(res, root.Val)

		rightArr := postorderTraversal(root.Right)
		res = reflect.AppendSlice(reflect.ValueOf(res), reflect.ValueOf(rightArr)).Interface().([]int)
	}
	return res
}

//迭代方式
func postorderTraversal_2(root *util.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := &util.Stack{}
	stack.Push(root)
	for !stack.Empty() {
		node := stack.Top().(*util.TreeNode)

		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			stack.Pop()
		}

		if node.Right != nil {
			stack.Push(node.Right)
			node.Right = nil
		}
		if node.Left != nil {
			stack.Push(node.Left)
			node.Left = nil
		}
	}
	return res
}

type op struct {
	s    string
	node *util.TreeNode
}

//使用额外结构体
func postorderTraversal_3(root *util.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	stack := &util.Stack{}
	stack.Push(&op{s: "go", node: root})

	for !stack.Empty() {
		top := stack.Top().(*op)
		stack.Pop()
		if top.s == "print" {
			res = append(res, top.node.Val)
			continue
		}

		stack.Push(&op{s: "print", node: top.node})
		if top.node.Right != nil {
			stack.Push(&op{s: "go", node: top.node.Right})
		}
		if top.node.Left != nil {
			stack.Push(&op{s: "go", node: top.node.Left})
		}
	}
	return res
}
