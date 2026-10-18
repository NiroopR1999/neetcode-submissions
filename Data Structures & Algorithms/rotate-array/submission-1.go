func rotate(nums []int, k int) {
	k=k%len(nums)
	left,right:=0,len(nums)-1

	for left<right {
		nums[left],nums[right]=nums[right],nums[left]
		left++
		right--
	}

	left,right=0,k-1
	for left<right {
		nums[left],nums[right]=nums[right],nums[left]
		left++
		right--
	}
	left,right=k,len(nums)-1
	for left<right {
			nums[left],nums[right]=nums[right],nums[left]
			left++
			right--
		}
}
