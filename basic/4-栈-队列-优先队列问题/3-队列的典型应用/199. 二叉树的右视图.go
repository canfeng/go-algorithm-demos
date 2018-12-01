package main

import (
	"algorithm-demo/util"
	"log"
)

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例:

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

	   1            <---
	 /   \
	2     3         <---
	 \     \
	  5     4       <---
	/
   7				<---
 */
func main() {
	res := rightSideView(util.ArrToTreeNode([]interface{}{1, 2, 3, nil, 5, nil, 4}))
	log.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *util.TreeNode) []int {
	//其实还是层序遍历，只是每层只取一个数，就是最后边的那个数
	var res []int
	if root == nil {
		return res
	}
	queue := []*Item{{order: 0, node: root}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1 : len(queue) : len(queue)]
		if item.node == nil {
			continue
		}
		order := item.order
		node := item.node
		if len(res) <= order {
			res = append(res, node.Val)
		}else{
			//每次都用最新的node替换同级的node
			res[order] = node.Val
		}
		queue = append(queue, &Item{order: order + 1, node: node.Left})
		queue = append(queue, &Item{order: order + 1, node: node.Right})
	}
	return res
}

type Item struct {
	order int //表示层级，从0开始，root为0级，root的子节点为1级
	node  *util.TreeNode
}
