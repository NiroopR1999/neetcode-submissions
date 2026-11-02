func groupAnagrams(strs []string) [][]string {
	group:=make(map[[26]int][]string)

	for i:=range strs {
		temp:=[26]int{}
		for j:=range strs[i]{
			temp[strs[i][j]-'a']++
		}
		group[temp]=append(group[temp],strs[i])
	}

	res:=[][]string{}

	for _,value:=range group{
		res=append(res,value)
	}
	return res
}
