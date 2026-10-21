func mergeAlternately(word1 string, word2 string) string {
	res:=[]byte{}
	left,right:=0,0

	for left<len(word1) && right<len(word2) {
		res=append(res,word1[left])
		res=append(res,word2[right])
		left++
		right++
	}
	
	for left<len(word1) {
		res=append(res,word1[left])
		left++
	}
	for right<len(word2) {
		res=append(res,word2[right])
		right++
	}
	return string(res)
}	
