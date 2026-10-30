/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    res:=0
	var InOrder func(*TreeNode)
	InOrder=func (node *TreeNode) {
		if node==nil || k==0 {
			return 
		}
		InOrder(node.Left)
		k--
		if k==0 {
			res=node.Val
			return
		}
		InOrder(node.Right)
	}
	InOrder(root)
	return res
}
