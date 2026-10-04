func majorityElement(nums []int) int {
    candidate,count:=0,0

	for i:=range nums {
		if count==0 {
			candidate=nums[i]
		}
		if nums[i]==candidate{
			count++
		} else {
			count--
		}
	}
	return candidate
}
