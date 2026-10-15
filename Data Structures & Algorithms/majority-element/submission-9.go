func majorityElement(nums []int) int {
    major,count:=0,0

	for i:=0;i<len(nums);i++ {
		if count==0 {
			major=nums[i]
			count++
		}
		if nums[i]!= major{
			count--
		} else {
			count++
		}
	}

	return major
}
