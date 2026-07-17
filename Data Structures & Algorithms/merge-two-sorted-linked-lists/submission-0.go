/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var head *ListNode
	var tail *ListNode // pointer of the new LL
	curr1 := list1 // pointer of LL1
	curr2 := list2 // pointer of LL2
	// defining the head
	if curr1.Val < curr2.Val { 
		head = curr1
		tail = curr1
		curr1 = curr1.Next
	} else {
		head = curr2
		tail = curr2
		curr2 = curr2.Next
	}
    for curr1 != nil && curr2 != nil {
		fmt.Printf("start: curr1: %d, curr2: %d, tail: %d\n", curr1.Val, curr2.Val, tail.Val)
		if curr1.Val < curr2.Val {
			tail.Next = curr1
			curr1 = curr1.Next
		} else {
			tail.Next = curr2
			curr2 = curr2.Next
		}
		tail = tail.Next
		fmt.Printf("tail val %d next %v\n", tail.Val, tail.Next)
	}
	tail.Next = curr1
	if curr1 == nil {
		tail.Next = curr2
	}
	return head
}
