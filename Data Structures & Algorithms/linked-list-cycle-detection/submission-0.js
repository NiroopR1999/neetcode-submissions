/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

class Solution {
    /**
     * @param {ListNode} head
     * @return {boolean}
     */
    hasCycle(head) {
            let current=head
        while(current)
        {
            if(current.visited) return true
            current.visited=true
            current=current.next
        }

        return false
    }
}
