func minSubArrayLen(target int, nums []int) int {
	total,minLength,left,right:=0,len(nums)+1,0,0

	for right < len(nums) {
		total+=nums[right]

		for total>=target {
			minLength=min(right-left+1,minLength)
			total-=nums[left]
			left+=1
		}
		right+=1
	}
	if minLength==len(nums)+1 {
		return 0
	}
	return minLength
}
