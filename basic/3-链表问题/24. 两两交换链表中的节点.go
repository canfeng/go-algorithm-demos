package __链表问题

import "algorithm-demo/util"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
说明:

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func main() {
	node := util.ArrToListNode(1, 2, 3, 4)
	node.Print()
	swapPairs(node).Print()
}

func swapPairs(head *util.ListNode) *util.ListNode {
	//设置虚拟头结点
	dummyHead := &util.ListNode{Val: 1, Next: head}

	pre := dummyHead
	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := node1.Next
		next := node2.Next
		//交换节点
		node2.Next = node1
		node1.Next = next
		pre.Next = node2

		//移到下一组节点
		pre = node1
	}
	return dummyHead.Next
}
