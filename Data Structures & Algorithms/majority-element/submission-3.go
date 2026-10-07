func majorityElement(nums []int) int {
    count:=0
	major:=0

	for i:=0;i<len(nums);i++ {
		if count==0 {
			major=nums[i]
		}
		if nums[i]==major{
			count++
		} else {
			count--
		}
	}
	return major
}
