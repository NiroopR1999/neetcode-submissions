func longestConsecutive(nums []int) int {
	seen:=make(map[int]bool)
	longest:=0
	for _,num:=range nums {
		seen[num]=true
	}

	for _,num:=range nums {
		if !seen[num-1] {
			temp:=num
			count:=0
			for seen[temp] {
				count++
				temp++
			}
			longest=max(longest,count)
		} 
	}

	return longest
}
