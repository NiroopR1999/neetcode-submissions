func longestConsecutive(nums []int) int {
	seen:=make(map[int]bool)
	longest:=0
	for _, num :=range nums {
		seen[num]=true
	}

	for  key, _ :=range seen {
		temp:=key
		if !seen[temp-1] {
			count:=0
			for seen[temp] {
				count+=1
				temp+=1
			}
			longest=max(longest,count)
		}
	}
	return longest
}
