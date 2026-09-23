func rotate(nums []int, k int) {
	k=k%len(nums)
	if len(nums)==k {
		return
	}
	left,right:=0,len(nums)-1

	for left<right {
		nums[left],nums[right]=nums[right],nums[left]
		right-=1
		left+=1
	}
	left,right=0,k-1
	for left<right {
		nums[left],nums[right]=nums[right],nums[left]
		right-=1
		left+=1
	}
	left,right=k,len(nums)-1
	for left<right {
		nums[left],nums[right]=nums[right],nums[left]
		right-=1
		left+=1
	}
}
