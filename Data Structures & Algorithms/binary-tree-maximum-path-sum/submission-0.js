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
    maxPathSum(root) {
        if(root===null) return 0
        let MPS=-Infinity

        const maxGain=(current)=>
        {
            if(current===null) return 0
            const leftGain=Math.max(maxGain(current.left),0)
            const rightGain=Math.max(maxGain(current.right),0)
            const CPS=current.val+leftGain+rightGain
            MPS=Math.max(CPS,MPS)
            return current.val+Math.max(leftGain,rightGain)
        }
        maxGain(root)
        return MPS
    }
}
