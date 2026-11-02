func longestCommonPrefix(strs []string) string {
    if len(strs)==0 {
		return ""
	}
	res:=strs[0]

	for i:=1;i<len(strs);i++  {
		j:=0
		for j<len(res) && j<len(strs[i]) {
			if res[j]!=strs[i][j] {
				break
			}
			j++
		}
		res=res[:j]
	}

	return res
}
