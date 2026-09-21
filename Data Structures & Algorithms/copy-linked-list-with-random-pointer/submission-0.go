/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head==nil {
		return head
	}
	copyMap:=make(map[*Node]*Node)

	current:=head
	
	for current!=nil {
	copyMap[current]=&Node{Val:current.Val}
	current=current.Next
	}

	current=head

	for current!=nil {
		copyNode:=copyMap[current]
		copyNode.Next=copyMap[current.Next]
		copyNode.Random=copyMap[current.Random]
		current=current.Next
	}

	return copyMap[head]
}
