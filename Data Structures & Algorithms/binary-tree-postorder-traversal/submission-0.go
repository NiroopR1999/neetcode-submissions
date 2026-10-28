/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res:=[]int{}
	return postOrder(root,res)
}

func postOrder(root *TreeNode, res []int) []int {
	if root==nil {
		return res
	}
	res=postOrder(root.Left,res)
	res=postOrder(root.Right,res)
	res=append(res,root.Val)
	return res
}