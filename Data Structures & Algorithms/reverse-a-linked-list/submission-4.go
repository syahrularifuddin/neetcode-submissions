/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, next *ListNode
	curr := head
    for {
		if curr.Next == nil {
			curr.Next = prev
			break
		}
		next = curr.Next
		if prev == nil {
			curr.Next = nil
		} else {
			curr.Next = prev
		}
		prev = curr
		curr = next
	}
	return curr
}
