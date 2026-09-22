/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy:=&ListNode{Next:head}
	leftPrev:=dummy

	for i:=1;i<left;i++ {
		leftPrev=leftPrev.Next
	}

	current:=leftPrev.Next

	var prev *ListNode

	for i:=1;i<=right-left+1;i++ {
		next:=current.Next
		current.Next=prev
		prev=current
		current=next
	}
	leftPrev.Next.Next=current
	leftPrev.Next=prev


	return dummy.Next
}
