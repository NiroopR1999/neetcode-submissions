func lengthOfLongestSubstring(s string) int {
	seen:=make(map[byte]bool)
	left,right,longest:=0,0,0

	for right<len(s) {
		if seen[s[right]] {
			seen[s[left]]=false
			left++
			continue
		}
		seen[s[right]]=true 
		longest=max(longest,right-left+1)

		right++
	}

	return longest
}
