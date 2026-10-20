/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	current:=head
	var prev *ListNode

	for current!=nil {
		temp:=current.Next
		current.Next=prev
		prev=current
		current=temp
	}
	head=prev

	return prev
}
