func minSubArrayLen(target int, nums []int) int {
	minLen:=len(nums)+1
	left:=0
	right:=0
	sum:=0
	for right<len(nums) {
		sum+=nums[right]
		for sum>=target{
			minLen=min(minLen,right-left+1)
			sum-=nums[left]
			left++
		}
		right++
	}
	if minLen ==len(nums)+1 {
		return 0
	}
	return minLen
}
