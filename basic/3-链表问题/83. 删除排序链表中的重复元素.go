package __链表问题

import (
	"algorithm-demo/util"
)

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */
func main() {
	deleteDuplicates(util.ArrToListNode()).Print()
	deleteDuplicates(util.ArrToListNode(1, 1, 2)).Print()
	deleteDuplicates(util.ArrToListNode(1, 1, 1)).Print()
	deleteDuplicates(util.ArrToListNode(1, 1, 2, 3, 3)).Print()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *util.ListNode) *util.ListNode {
	if head == nil {
		return nil
	}

	//排序数组说明重复的元素肯定是相邻的
	var dummyHead = &util.ListNode{Val: 0, Next: head}
	cur := head
	for cur.Next != nil {
		next := cur.Next
		if next.Val != cur.Val {
			cur = next
		} else {
			cur.Next = next.Next
		}
	}
	return dummyHead.Next
}
