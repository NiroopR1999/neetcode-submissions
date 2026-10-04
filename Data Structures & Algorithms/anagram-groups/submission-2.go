func groupAnagrams(strs []string) [][]string {
	groups:=make(map[[26]int][]string)

	for _,str:=range strs {
		var key [26]int
		for _, char :=range str {
			key[char-'a']=key[char-'a']+1
		}

		groups[key]=append(groups[key],str)
	}
	res:=make([][]string,0,len(strs))

	for _,val:=range groups {
		res=append(res,val)
	}
	return res
}
