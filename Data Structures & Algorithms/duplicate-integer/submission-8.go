func hasDuplicate(nums []int) bool {
	seen:=make(map[int]int)
    for i:=range nums{
		seen[nums[i]]++
		if seen[nums[i]]>1 {
			return true
		}
	}
	return false
}
