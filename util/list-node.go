package util

import "encoding/json"

func ArrToListNode(arr ...int) *ListNode {
	var ret *ListNode
	var next *ListNode = nil
	for i := len(arr) - 1; i >= 0; i-- {
		ret = &ListNode{arr[i], next}
		next = ret
	}
	return ret
}

func (node *ListNode) ToArr() []int {
	var out []int
	for node != nil {
		out = append(out, node.Val)
		node = node.Next
	}
	return out
}

func (head *ListNode) Print() {
	cur := head
	for cur != nil {
		print(cur.Val)
		print("->")
		cur = cur.Next
	}
	print("nil\n")
}

func (node *ListNode) ToString() string {
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
