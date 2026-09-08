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
     * @return {number[][]}
     */
    levelOrder(root) {

        if(root===null) return []
        const queue= new Queue()
        const res=[]
        queue.enqueue(root)

        while(!queue.isEmpty())
        {
            const len=queue.size()
            const nodes=[]

            for(let i=0;i<len;i++)
            {
                const current=queue.dequeue()
                nodes.push(current.val)
                if(current.left)
                {
                    queue.enqueue(current.left)
                }

                if(current.right)
                {
                    queue.enqueue(current.right)
                }
            }
            res.push(nodes)
        }
        return res
    }
}
