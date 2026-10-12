func productExceptSelf(nums []int) []int {
	prefix,suffix:=1,1
	res:=make([]int,len(nums))

	for i:=range nums {
		res[i]=prefix
		prefix*=nums[i]
	}

	for i:=len(nums)-1;i>=0;i-- {
		res[i]*=suffix
		suffix*=nums[i]
	}

	return res
}
