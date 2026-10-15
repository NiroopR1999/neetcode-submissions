func longestCommonPrefix(strs []string) string {
    if len(strs)==0 {
		return ""
	}
	longest:=strs[0]

	for i:=1;i<len(strs);i++ {
		j:=0
		for j<len(longest) && j<len(strs[i]) {
			if longest[j]!=strs[i][j]{
				break
			}
			j++
		}
		longest=longest[:j]
	}

	return longest
}
