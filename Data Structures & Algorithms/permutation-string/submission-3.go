func checkInclusion(s1 string, s2 string) bool {
	
	if len(s2)<len(s1) {
		return false
	}

	freq1:=[26]int{}
	freq2:=[26]int{}

	for i:=range s1 {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}
	if freq1==freq2 {
		return true
	}
	left,right:=0,len(s1)
	for right<len(s2) {
		freq2[s2[left]-'a']--
		freq2[s2[right]-'a']++
		left++
		right++
		if freq1==freq2 {
			return true
		}
	}

	return false
}
