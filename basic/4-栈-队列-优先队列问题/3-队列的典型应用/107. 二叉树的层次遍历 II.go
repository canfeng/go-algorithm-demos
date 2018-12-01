package main

import (
	"algorithm-demo/util"
	"log"
)

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]
 */
func main() {
	bottom := levelOrderBottom(util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}))
	log.Println(bottom)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *util.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*Item{{order: 0, node: root}}

	for len(queue) > 0 {
		item := queue[0]
		//移除第一个元素
		queue = queue[1:len(queue):len(queue)]

		if len(res) <= item.order || res[item.order] == nil {
			res = append(res, []int{})
		}

		res[item.order] = append(res[item.order], item.node.Val)

		if item.node.Left != nil {
			queue = append(queue, &Item{order: item.order + 1, node: item.node.Left})
		}
		if item.node.Right != nil {
			queue = append(queue, &Item{order: item.order + 1, node: item.node.Right})
		}
	}
	//反转数组
	len := len(res)
	for i := 0; i < len/2; i++ {
		res[i], res[len-1-i] = res[len-1-i], res[i]
	}

	return res
}

type Item struct {
	order int //表示层级，从0开始，root为0级，root的子节点为1级
	node  *util.TreeNode
}
