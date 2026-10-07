func hasDuplicate(nums []int) bool {
    seen:=make(map[int]int)

	for _,num:=range nums {
		seen[num]++
		if seen[num] > 1 {
			return true
		}
	}
	return false
}
