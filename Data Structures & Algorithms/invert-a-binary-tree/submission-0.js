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
     * @return {TreeNode}
     */
    invertTree(root) {

        const invertBT=(currentRoot=root)=>
        {
                if(currentRoot)
                {
                    invertBT(currentRoot.left)
                    invertBT(currentRoot.right)
                    const temp=currentRoot.left
                    currentRoot.left=currentRoot.right
                    currentRoot.right=temp
                }
        }
        invertBT()
        return root
    }
}
