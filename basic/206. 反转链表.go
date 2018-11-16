package main

import "encoding/json"

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 */
func main() {
	list := reverseList_iteration(ArrToListNode(1, 2, 3, 4, 5))
	println(list.toString())
	list1 := reverseList_recursion(ArrToListNode(1, 2, 3, 4, 5))
	println(list1.toString())
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//迭代方式
func reverseList_iteration(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//递归方式
func reverseList_recursion(head *ListNode) *ListNode {
	return recuresionListNode(head, nil)
}

func recuresionListNode(node *ListNode, last *ListNode) *ListNode{
	if node != nil {
		next := node.Next
		node.Next = last
		last = node
		return recuresionListNode(next, last)
	}else{
		return last
	}
}

func ArrToListNode(arr ...int) *ListNode {
	var ret *ListNode
	var next *ListNode = nil
	for i := len(arr) - 1; i >= 0; i-- {
		ret = &ListNode{arr[i], next}
		next = ret
	}
	return ret
}

func (node *ListNode) toArr() []int {
	var out []int
	for node != nil {
		out = append(out, node.Val)
		node = node.Next
	}
	return out
}

func (node *ListNode) toString() string {
	bytes, err := json.Marshal(node)
	if err != nil {
		return ""
	}
	return string(bytes)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
