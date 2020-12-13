package dynamic

//K 个一组翻转链表
//给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
//
//k是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序
//
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverseList(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return hair.Next
}
func reverseList(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var prev *ListNode
	var curr *ListNode
	prev = tail.Next
	curr = head
	for prev != tail {
		var tempNext *ListNode
		tempNext = curr.Next
		curr.Next = prev
		prev = curr
		curr = tempNext
	}
	return tail, head
}
