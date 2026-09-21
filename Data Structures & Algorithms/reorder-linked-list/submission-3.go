/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    
	if head==nil || head.Next==nil{
		return
	}

	slow,fast:=head, head

	for fast!=nil && fast.Next !=nil {
		slow=slow.Next
		fast=fast.Next.Next
	}
	current:=slow.Next
	slow.Next=nil
	var prev *ListNode
	for current !=nil {
		temp:=current.Next
		current.Next=prev
		prev=current
		current=temp
	}
	first:=head
	second:=prev

	for second!=nil {
		firstNext:=first.Next
		secondNext:=second.Next

		first.Next=second
		second.Next=firstNext

		first=firstNext
		second=secondNext
	}


}
