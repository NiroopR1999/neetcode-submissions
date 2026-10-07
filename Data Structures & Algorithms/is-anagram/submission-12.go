func isAnagram(s string, t string) bool {

	if len(s)!=len(t) {
		return false
	}
	freq:=make(map[byte]int)

	for i :=range s{
		freq[t[i]]++
		freq[s[i]]--
	}	

	for _,val:=range freq {
		if val>0 {
			return false
		}
	}

	return true
}
