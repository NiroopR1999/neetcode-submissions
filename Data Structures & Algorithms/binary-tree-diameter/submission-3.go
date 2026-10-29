/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    dia:=0
	depth(root,&dia)
	return dia
}

func depth(root *TreeNode, dia *int) int {
	if root==nil {
		return 0
	}
	left:=depth(root.Left,dia)
	right:=depth(root.Right,dia)
	*dia=max(*dia,left+right)
	return max(left,right)+1
}