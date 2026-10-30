/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return check(root, math.MinInt64,math.MaxInt64)
}

func check(root *TreeNode, min, max int64) bool {
	if root==nil {
		return true
	}

	if int64(root.Val)<=min || int64(root.Val)>=max {
		return false
	}

	return check(root.Left, min, int64(root.Val)) && check(root.Right, int64(root.Val), max)
}