func isAnagram(s string, t string) bool {
	if len(s)!=len(t) {
		return false
	}
	a1,a2:=[26]int{},[26]int{}

	for _,char:=range s{
		a1[char-'a']++
	}
		for _,char:=range t{
		a2[char-'a']++
	}
	return a1==a2
}

