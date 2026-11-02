func minSubArrayLen(target int, nums []int) int {
	res:=len(nums)+1
	left,right:=0,0
	sum:=0
	for right<len(nums) {
		sum+=nums[right]
		for sum>=target {
			res=min(res,right-left+1)
			sum-=nums[left]
			left++
		}
		right++
	}
	if res==len(nums)+1 {
		return 0
	}
	return res

}
