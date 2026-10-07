func productExceptSelf(nums []int) []int {
	res:=make([]int,len(nums))
	prefix,suffix:=1,1

	for index,num:= range nums {
		res[index]=prefix
		prefix*=num
	}
	for i:=len(nums)-1;i>=0;i--{
		res[i]=res[i]*suffix
		suffix*=nums[i]
		
	}
	return res
}
