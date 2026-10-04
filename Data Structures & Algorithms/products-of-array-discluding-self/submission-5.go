func productExceptSelf(nums []int) []int {

	res:=make([]int,len(nums))
	prefix,suffix:=1,1
	for i:=range nums{
		res[i]=prefix
		prefix*=nums[i]
	}
	i:=len(nums)-1
	for i>=0{
		res[i]*=suffix
		suffix*=nums[i]
		i--
	}
	return res
}
