func containsNearbyDuplicate(nums []int, k int) bool {
	seen:=make(map[int]int)

	for i, num:=range nums{
		prevIndex, exist := seen[num] ; if exist {
			if i-prevIndex<=k {
				return true
			}
		}
		seen[num]=i
	}
	return false
}
