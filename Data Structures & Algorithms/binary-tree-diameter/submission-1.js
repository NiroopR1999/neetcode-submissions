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
    diameterOfBinaryTree(root) {

        let max=0

        const maximumDiameter=(currentRoot=root)=>
        {
            if(!currentRoot) return 0
                const leftMax=maximumDiameter(currentRoot.left) 
                const rightMax=maximumDiameter(currentRoot.right)
                max=Math.max(max,leftMax+rightMax)
                return 1+Math.max(leftMax,rightMax)
        }
        maximumDiameter()
        return max
    }
}
