package __链表问题

import "algorithm-demo/util"

/**
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
 */
func main() {
	n1 := util.ArrToListNode(1, 1, 2, 6, 3, 4, 5, 6)
	n1.Print()
	removeElements(n1, 1).Print()
	removeElements(nil, 1).Print()
	removeElements(nil, 1).Print()
	n2 := util.ArrToListNode(1, 1, 1, 1)
	n2.Print()
	removeElements(n2, 1).Print()
}

/**
 传统解法
 */
func removeElements_old(head *util.ListNode, val int) *util.ListNode {
	//先检查头结点（循环是针对于头部连续都是目标值的情况）
	for head != nil && head.Val == val {
		head = head.Next
	}
	//该判断是针对整个链表都是目标值
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			//如果为目标值，则更新next，cur不动
			cur.Next = next.Next
		} else {
			//否则，cur后移一位
			cur = cur.Next
		}
	}
	return head
}

/**
 使用虚拟头结点的解法
 */
func removeElements(head *util.ListNode, val int) *util.ListNode {
	//添加一个虚拟头结点，next指向head
	dummyNode := &util.ListNode{Val: 1, Next: head}

	cur := dummyNode
	for cur.Next != nil {
		next := cur.Next
		if next.Val == val {
			//如果为目标值，则更新next，cur不动
			cur.Next = next.Next
		} else {
			//否则，cur后移一位
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
