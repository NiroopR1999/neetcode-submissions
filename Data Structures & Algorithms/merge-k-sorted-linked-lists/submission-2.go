/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
	APPROACH:

	We have K sorted linked lists.

	Instead of repeatedly merging complete lists,
	we keep only the CURRENT smallest node from
	each list inside a min heap.

	Steps:

	1. Put the first node of every non-empty list
	   into the min heap.

	2. Take the smallest node from the heap.

	3. Add that node to our result.

	4. If that node has a next node, put that next
	   node into the heap.

	5. Repeat until the heap is empty.

	Why does this work?

	Because every individual list is already sorted.

	If we take a node from a list:

	    1 → 4 → 5
	    ↑

	after taking 1, only 4 can become the next
	candidate from that list.

	Time:  O(N log K)
	Space: O(K)

	N = total number of nodes
	K = number of lists
*/

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

// Smaller value should come out first.
// Therefore this is a MIN heap.
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {

	old := *h
	n := len(old)

	// container/heap has already moved the
	// smallest element to the end.
	node := old[n-1]

	// Remove it from the heap.
	*h = old[:n-1]

	return node
}

func mergeKLists(lists []*ListNode) *ListNode {

	// No lists.
	if len(lists) == 0 {
		return nil
	}

	// Create an empty min heap.
	h := &MinHeap{}

	heap.Init(h)

	// Put the FIRST node of every list
	// into the heap.
	for _, list := range lists {

		// Ignore empty lists.
		if list != nil {
			heap.Push(h, list)
		}
	}

	// Dummy node makes building the result easier.
	dummy := &ListNode{}

	// tail always points to the last node
	// in our result.
	tail := dummy

	for h.Len() > 0 {

		// Get the smallest current node.
		node := heap.Pop(h).(*ListNode)

		// Connect it to our result.
		tail.Next = node

		// Move tail forward.
		tail = tail.Next

		// node came from one particular list.
		//
		// Its next node is the next candidate
		// from that list.
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	// dummy itself is not part of the answer.
	return dummy.Next
}