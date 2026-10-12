func longestConsecutive(nums []int) int {
	seen:=make(map[int]bool)

	for _,num :=range nums {
		seen[num]=true
	}

	longest:=0

	for num,_:=range seen {
		if !seen[num-1] {
			count:=0
			for seen[num+count] {
				count++
			}
			longest=max(longest,count)
		}
	}
	return longest
}	
