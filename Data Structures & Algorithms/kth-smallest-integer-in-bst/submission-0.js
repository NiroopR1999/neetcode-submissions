/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     constructor(val = 0, left = null, right = null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    /**
     * @param {TreeNode} root
     * @param {number} k
     * @return {number}
     */
    kthSmallest(root, k) {

        if(root===null) return null
        const res=[]

        function inOrder(node=root)
        {
            if(node===null) return 
            inOrder(node.left)
            res.push(node.val)
            inOrder(node.right)
        }
        inOrder()

        return res[k-1]
    }
}
