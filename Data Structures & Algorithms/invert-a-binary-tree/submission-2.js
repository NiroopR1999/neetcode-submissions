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

        function invert(current=root)
        {  
            if(!current) return
            if(current)
            {   
                invert(current.left)
                invert(current.right) 
                const temp=current.left
                current.left=current.right
                current.right=temp
            }
        }
        invert()
        return root
    }
}
