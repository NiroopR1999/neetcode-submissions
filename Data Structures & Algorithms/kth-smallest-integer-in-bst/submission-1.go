/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	if root==nil {
		return 0
	}
    res:=[]int{}
	var InOrder func(*TreeNode)
	InOrder=func (node *TreeNode) {
		if node==nil {
			return 
		}
		InOrder(node.Left)
		res=append(res,node.Val)
		InOrder(node.Right)
	}
	InOrder(root)

	return res[k-1]
}
