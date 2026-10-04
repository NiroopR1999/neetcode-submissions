func isAnagram(s string, t string) bool {

	if len(s)!=len(t) {
		return false
	}
	res:=make(map[byte]int)
	for i:=range s {
		res[s[i]]++
		res[t[i]]--
	}
	for _,value :=range res{
		if value !=0 {
			return false
		}
	}
	return true
}
