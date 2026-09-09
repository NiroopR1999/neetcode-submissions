func isAnagram(s string, t string) bool {
	res:=make(map[rune]int)

	for _,v:= range s {
		res[v]=res[v]+1
	}
	for _,v:= range t {
		res[v]=res[v]-1
	}
	for _,v:= range res {
		if v!=0 {
			return false
		}
	}
	return true
}
