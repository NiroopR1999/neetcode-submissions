func longestCommonPrefix(strs []string) string {
    
	prefix:=strs[0]
	for i:=1;i<len(strs);i++ {
		 j:=0
		for j < len(strs[i]) && j < len(prefix) {
			if prefix[j]!=strs[i][j] {
				break
			} 
			j++
		}
		prefix=prefix[:j]
	}

	return prefix
}
