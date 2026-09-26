func maxSlidingWindow(nums []int, k int) []int {
    n,left,right:=len(nums),0,0
	q,res:=[]int{},[]int{}

	for right < n {
		if len(q)>0 && q[0] <left {
			q=q[1:]
		}

		for len(q) > 0 && nums[q[len(q)-1]] < nums[right] {
			q=q[:len(q)-1]
		}

		q=append(q,right)

		if right-left+1==k{
			res=append(res,nums[q[0]])
			left+=1
		}
		right+=1
	}
	return res
}
