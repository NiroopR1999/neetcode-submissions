func maxSlidingWindow(nums []int, k int) []int {
    queue:=[]int{}
	res:=[]int{}
	for i:=0;i<k;i++ {
		for len(queue)>0 &&  nums[queue[len(queue)-1]]<=nums[i]{
			queue=queue[:len(queue)-1]
		}
		queue=append(queue,i)
	}
	res=append(res,nums[queue[0]])
	left,right:=1,k

	for right<len(nums) {
		if queue[0]<left {
			queue=queue[1:]
		}
		for len(queue)>0 && nums[queue[len(queue)-1]]<=nums[right]{
			queue=queue[:len(queue)-1]
		}
		queue=append(queue,right)
		res=append(res,nums[queue[0]])
		left++
		right++
	}
	return res
}
