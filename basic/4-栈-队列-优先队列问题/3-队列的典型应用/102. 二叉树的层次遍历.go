package main

import (
	"algorithm-demo/util"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/pkg/errors"
	"log"
)

/**
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/
func main() {
	root := util.ArrToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	res := levelOrder(root)
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
/**
使用队列的方式保存层级的顺序
*/
func levelOrder(root *util.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var que *queue.Queue
	que.Put(&Item{order: 0, node: root})

	for !que.Empty() {
		//取出队列的第一个元素
		peek, err := que.Peek()
		if err != nil {
			errors.Wrap(err, "queue.peek() error")
		}
		one := peek.(*Item)

		if res == nil || len(res) <= one.order || res[one.order] == nil {
			//如果还不存在同一层级的数组，则新创建一个
			res = append(res, []int{})
		}
		res[one.order] = append(res[one.order], one.node.Val)

		//将左右节点分别推入队列
		if one.node.Left != nil {
			que.Put(one.node.Left)
		}
		if one.node.Right != nil {
			que.Put(one.node.Right)
		}
	}

	return res
}

type Item struct {
	order int //表示层级，从0开始，root为0级，root的子节点为1级
	node  *util.TreeNode
}

func levelOrder_2(root *util.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*Item
	queue = append(queue, &Item{order: 0, node: root})

	for len(queue) > 0 {
		//取出队列的第一个元素
		one := queue[0]
		//移除第一个元素
		queue = queue[1:len(queue):len(queue)]

		if res == nil || len(res) <= one.order || res[one.order] == nil {
			//如果还不存在同一层级的数组，则新创建一个
			res = append(res, []int{})
		}
		res[one.order] = append(res[one.order], one.node.Val)

		//将左右节点分别推入队列
		if one.node.Left != nil {
			queue = append(queue, &Item{order: one.order + 1, node: one.node.Left})
		}
		if one.node.Right != nil {
			queue = append(queue, &Item{order: one.order + 1, node: one.node.Right})
		}
	}

	return res
}
