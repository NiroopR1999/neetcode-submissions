func lengthOfLongestSubstring(s string) int {
	maxLen,left,right:=0,0,0
	seen:=make(map[byte]int)

	for right<len(s) {
		index,exists:=seen[s[right]]
		
		if exists{
			left=max(index+1,left)
		}
		seen[s[right]]=right
		maxLen=max(maxLen,right-left+1)
		right+=1
	}
	return maxLen
}
