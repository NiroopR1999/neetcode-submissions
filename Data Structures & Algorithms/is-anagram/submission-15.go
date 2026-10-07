func isAnagram(s string, t string) bool {

	if len(s)!=len(t) {
		return false
	}
	first,second:=[26]int{},[26]int{}

	for i:=range s {
		first[s[i]-'a']++
		second[t[i]-'a']++
	}
	return first==second
}
