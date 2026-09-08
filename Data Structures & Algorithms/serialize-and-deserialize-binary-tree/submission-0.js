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

class Codec {
    /**
     * Encodes a tree to a single string.
     *
     * @param {TreeNode} root
     * @return {string}
     */
    serialize(root) {
        
        if(root===null) return ''
        const res=[]

        const dfs=(current)=>
        {
            if(current===null)
            {
               return res.push('null')
                
            }
            res.push(current.val)
            dfs(current.left)
            dfs(current.right)
        }
        dfs(root)
        return res.join(',')
    }

    /**
     * Decodes your encoded data to tree.
     *
     * @param {string} data
     * @return {TreeNode}
     */
    deserialize(data) {
         if (data === '') return null;  // Handle the edge case of an empty tree
       let i=0
       const bt=data.split(',')
       const dfs=()=>
       {
        if(bt[i]==='null'){
            i++
            return null
        }
        const node= new TreeNode(parseInt(bt[i++]))
        node.left=dfs()
        node.right=dfs()
        return node
       }
       return dfs()
    }
}
