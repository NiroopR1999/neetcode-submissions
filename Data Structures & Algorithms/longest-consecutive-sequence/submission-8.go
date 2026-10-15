func longestConsecutive(nums []int) int {
	seen:=make(map[int]bool)
	longest:=0
	for _,num:=range nums{
		seen[num]=true
	}

	for _,num:=range nums{
		if !seen[num-1]{
			count:=0
			for seen[num+count] {
				count++
			}
			longest=max(count,longest)
		}
	}
	return longest
}
