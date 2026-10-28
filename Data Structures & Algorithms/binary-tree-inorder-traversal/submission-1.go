/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	res:=[]int{}
	return inorder(root, res)
}

func inorder(root *TreeNode, res []int) []int {
	if root==nil {
		return res
	}
	res=inorder(root.Left, res)
	res=append(res,root.Val)
	res=inorder(root.Right, res)
	return res
}
