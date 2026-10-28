/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res:=[][]int{}
	queue:=[]*TreeNode{}
	if root==nil {
		return res
	}
	queue=append(queue,root)

	for len(queue)>0 {
		levelLen:=len(queue)
		level:=[]int{}
		for i:=0;i<levelLen;i++ {
			pop:=queue[0]
			queue=queue[1:]
			level=append(level,pop.Val)
			if pop.Left!=nil {
			queue=append(queue,pop.Left)
			}
			if pop.Right!=nil {
			queue=append(queue,pop.Right)
			}
		}
		res=append(res,level)
	}
	return res
}
