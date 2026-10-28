/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    res:=[]int{}
	return preOrder(root,res)
}


func preOrder(root *TreeNode, res []int) []int {
	if root==nil {
		return res
	}
	res=append(res,root.Val)
	res=preOrder(root.Left, res)
	res=preOrder(root.Right, res)
	return res
}