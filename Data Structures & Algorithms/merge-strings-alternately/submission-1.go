func mergeAlternately(word1 string, word2 string) string {
	n1,n2:=len(word1),len(word2)
	res:=[]byte{}
	first,second:=0,0
	for first<n1 && second<n2 {
		res=append(res,word1[first])
		res=append(res,word2[second])
		first++
		second++
	}
	for first<n1 {
		res=append(res,word1[first])
		first++

	}
	for second<n2 {
		res=append(res,word2[second])
		second++
	}
	return string(res)

}
