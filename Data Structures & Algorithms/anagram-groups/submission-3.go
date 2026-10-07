func groupAnagrams(strs []string) [][]string {
	match:=make(map[[26]int][]string)

	for i:=range strs {
		key:=[26]int{}
		for j:=range strs[i] {
			key[strs[i][j]-'a']++
		}
		match[key]=append(match[key],strs[i])
	}
	res:=make([][]string,0,len(strs))
	for _,val:=range match {
		res=append(res,val)
	}
	return res
}
