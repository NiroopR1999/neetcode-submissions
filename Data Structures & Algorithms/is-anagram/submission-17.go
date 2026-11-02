func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	a1,a2:=[26]int{},[26]int{}

	for i:=range s {
		a1[s[i]-'a']++
		a2[t[i]-'a']++
	}

	return a1==a2
}
