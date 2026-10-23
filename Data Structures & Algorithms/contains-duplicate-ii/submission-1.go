func containsNearbyDuplicate(nums []int, k int) bool {
		seen:=make(map[int]bool)
		left,right:=0,0

		for right<len(nums) {
			if right-left>k {
				seen[nums[left]]=false
				left++
			}
			if seen[nums[right]]{
				return true
			}
			seen[nums[right]]=true
			right++
		}

	return false
}
