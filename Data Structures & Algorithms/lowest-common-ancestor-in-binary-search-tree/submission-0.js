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
     * @param {TreeNode} p
     * @param {TreeNode} q
     * @return {TreeNode}
     */
    lowestCommonAncestor(root, p, q) {

        let p1=[]
        let q1=[]
        function search(node)
        {   
            let path=[]
            let current=root

            while(current!==null)
            {
                path.push(current)
                if(current.val===node.val) break
                if(node.val<current.val)
                {
                    current=current.left
                }
                else{
                    current=current.right
                }
            }
            return path
        }
        p1=search(p)
        q1=search(q)
        let m=p1.length
        let n=q1.length
        let i=0
        let LCA=null
        while(i<m && i<n && p1[i].val===q1[i].val)
        {
            LCA=p1[i++]
        }
        return LCA
    }
}
