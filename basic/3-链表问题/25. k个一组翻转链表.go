package __链表问题

import "algorithm-demo/util"

/*
给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func main() {
	reverseKGroup(util.ArrToListNode(1, 2, 3, 4, 5), 5).Print()
}

func reverseKGroup(head *util.ListNode, k int) *util.ListNode {
	//采用从后往前的方式反转,先反转最后一组，将最后一组反转后的head返回，作为倒数第二组的第一个元素的next

	//先检查是否足够完整的一组
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	//递归，走下一组
	t := reverseKGroup(cur, k)

	//反转
	var c *util.ListNode
	for k > 0 {
		c = head
		head = head.Next
		c.Next = t
		t = c
		k--
	}

	return t
}
