package __链表问题

import (
	"algorithm-demo/util"
	"log"
	"math"
	"strconv"
)

//给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
//
//你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

func main() {
	l1 := util.ArrToListNode(1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1)
	arr1 := l1.ToArr()
	log.Println("l1 :", arr1, "; lens:", len(arr1))
	l2 := util.ArrToListNode(5, 6, 4)
	arr2 := l2.ToArr()
	log.Println("l2 :", arr2, "; lens:", len(arr2))
	ret := addTwoNumbers(l1, l2)
	log.Println("ret:", ret.ToArr())
}

//正确解法
func addTwoNumbers2(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	x, y, carry := 0, 0, 0
	var final = &util.ListNode{}
	var ret = final
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			x = l1.Val
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		v := sum % 10
		carry = sum / 10

		ret.Next = &util.ListNode{Val: v}
		ret = ret.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return final.Next
}

//会有整型溢出的情况
func addTwoNumbers(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {

	v1 := 0
	for i := 0; l1 != nil; i++ {
		v1 += l1.Val * int(math.Pow10(i))
		l1 = l1.Next
	}
	v2 := 0
	for i := 0; l2 != nil; i++ {
		v2 += l2.Val * int(math.Pow10(i))
		l2 = l2.Next
	}
	log.Println("v1:", v1)
	log.Println("v2:", v2)

	sum := int64(v1) + int64(v2)
	log.Println("sum:", sum)
	arr := make([]int, len(strconv.FormatInt(sum, 10)))
	for i := 0; sum > 0; i++ {
		val := sum % 10
		arr[i] = int(val)
		sum = sum / 10
	}

	return util.ArrToListNode(arr...)

}
