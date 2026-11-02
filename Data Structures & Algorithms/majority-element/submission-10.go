func majorityElement(nums []int) int {
    major:=nums[0]
	count:=1
	for i:=1;i<len(nums);i++{
		if count==0 {
			major=nums[i]
		} else if nums[i]!=major{
			count--
		} else {
			count++	
		}
	}
	return major
}
