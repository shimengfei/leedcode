package dynamic

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	prev = nil
	var p *ListNode
	p = head
	//p==nil才为终止条件,因为p.Next==nil此时p指向最后一个元素，需要将该元素翻转
	for p != nil {
		var tempNext *ListNode
		tempNext = p.Next
		p.Next = prev
		prev = p
		p = tempNext
	}
	return prev
}
