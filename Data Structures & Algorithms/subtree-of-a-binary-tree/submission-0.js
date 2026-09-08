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
     * @param {TreeNode} subRoot
     * @return {boolean}
     */
    isSubtree(root, subRoot) {

        function isSame(p,q)
        {
            if(p===null && q===null) return true

            if(p===null || q===null) return false

            if(p.val!==q.val) return false

            return isSame(p.left,q.left) && isSame(p.right,q.right)
        }

        function traverse(current=root)
        {
            if(current===null) return false
            if(current.val===subRoot.val && isSame(current,subRoot)) return true
            return traverse(current.left) || traverse(current.right)
        }   

        return traverse()
    }
}
