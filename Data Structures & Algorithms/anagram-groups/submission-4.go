func groupAnagrams(strs []string) [][]string {
	freq:=make(map[[26]int][]string)
	res:=[][]string{}

	for _,str:=range strs {
		arr:=[26]int{}
		for _,char:=range str{
			arr[char-'a']++
		}
		freq[arr]=append(freq[arr],str)
	}

	for _,value:=range freq{
		res=append(res,value)
	}
	return res
}
