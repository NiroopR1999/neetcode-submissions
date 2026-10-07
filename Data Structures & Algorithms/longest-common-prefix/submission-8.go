func longestCommonPrefix(strs []string) string {
    if len(strs)<=0{
		return ""
	}
	sample:=strs[0]

	for i:=1;i<len(strs);i++ {
		j:=0
		for j<len(sample) && j<len(strs[i]) && sample[j]==strs[i][j] {
			j++
		}
		sample=strs[i][0:j]
	}
	return sample
}
