package main

import (
	"algorithm-demo/util"
	"log"
)

/**
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
 */
func main() {
	res1 := zigzagLevelOrder(util.ArrToTreeNode([]interface{}{1, 2, 3, 4, nil, nil, 5}))
	log.Println(res1)
	res := zigzagLevelOrder(util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7}))
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
func zigzagLevelOrder(root *util.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*Item{{order: 0, node: root}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:len(queue):len(queue)]

		order := item.order
		node := item.node

		if node == nil {
			continue
		}
		log.Println("current item=", order, node.Val)

		if len(res) <= order {
			res = append(res, []int{})
		}
		res[order] = append(res[order], item.node.Val)

		nextOrder := order + 1
		queue = append(queue, &Item{order: nextOrder, node: node.Left})
		queue = append(queue, &Item{order: nextOrder, node: node.Right})
	}
	//反转偶数索引的数组
	for i := 0; i < len(res); i++ {
		if i%2 == 1 {
			util.ReverseArray(res[i])
		}
	}
	return res
}

type Item struct {
	order int //表示层级，从0开始，root为0级，root的子节点为1级
	node  *util.TreeNode
}
