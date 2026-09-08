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
     * @param {ListNode} list1
     * @param {ListNode} list2
     * @return {ListNode}
     */
    mergeTwoLists(list1, list2) {
        const res= new ListNode()
        let current=res
        while (list1 && list2 )
        {
            if(list1.val <= list2.val)
            {
                current.next=list1
                list1=list1.next
            }
            else{
                current.next=list2
                list2=list2.next
            }
            current=current.next
        }
        // Append remaining nodes of list1 or list2
        if (list1) {
            current.next = list1;
        } else {
            current.next = list2;
        }
        return res.next
        
    }
}
