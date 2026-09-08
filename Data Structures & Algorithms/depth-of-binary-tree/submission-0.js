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

        if (root === null) return 0

       const  maximunDepth= (currentRoot) => {
            if (currentRoot === null) return 0
            let maxLeft = maximunDepth(currentRoot.left)
            let maxRight = maximunDepth(currentRoot.right)
            return 1 + Math.max(maxLeft, maxRight)
        }

        return maximunDepth(root)
    }
}