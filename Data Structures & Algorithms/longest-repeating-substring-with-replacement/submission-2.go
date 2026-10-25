func characterReplacement(s string, k int) int {
	left,right,maxFreq,longest:=0,0,0,0
	freq:=[26]int{}

	for right<len(s) {
		freq[s[right]-'A']++
		maxFreq=max(maxFreq,freq[s[right]-'A'])
		windowLen:=right-left+1
		if windowLen-maxFreq>k{
			freq[s[left]-'A']--
			left++
		}
		windowLen=right-left+1
		longest=max(longest,windowLen)
		right++
	}
	return longest
}
