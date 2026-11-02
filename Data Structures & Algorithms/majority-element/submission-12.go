func majorityElement(nums []int) int {
    major:=0
	count:=0
	for i:=0;i<len(nums);i++{
		if count==0 {
			major=nums[i]
			count++
		} else if nums[i]!=major{
			count--
		} else {
			count++	
		}
	}
	return major
}
