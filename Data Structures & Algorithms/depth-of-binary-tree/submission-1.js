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
     * @return {number}
     */
    maxDepth(root) {

        const height = (current=root)=>
        {
            if(current===null) return 0

            const left=height(current.left)
            const right=height(current.right)

            return 1+Math.max(left,right)
        }

        return height()
    }
}
