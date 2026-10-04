func twoSum(nums []int, target int) []int {
    seen:=make(map[int]int) 
	for i:=range nums {
		if index,ok :=seen[target-nums[i]]; ok {
			return []int{index,i}
		} 
		seen[nums[i]]=i
	}
	return []int{}
}

