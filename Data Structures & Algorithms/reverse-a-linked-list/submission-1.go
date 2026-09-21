/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    
	if head==nil {
		return head
	}

	var prev *ListNode
	current:=head

	for current !=nil {
		temp:=current.Next
		current.Next=prev
		prev=current
		current=temp
	}

	head=prev

	return head
}
