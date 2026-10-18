func isPalindrome(s string) bool {
	newStr:=[]byte{}

	for i:=range s {
		if isAplhanumeric(s[i]) {
			newStr=append(newStr,toLower(s[i]))
		}
	}

	left,right:=0,len(newStr)-1
	for left<=right{
		if newStr[left]!=newStr[right]{
			return false
		}
		left++
		right--
	}
	return true
}

func isAplhanumeric(c byte) bool {
	return (c>='a' && c<='z') || (c>='A' && c<='Z') || (c>='0' && c<='9')
}

func toLower(c byte) byte {
	if (c>='A' && c<='Z') {
		return c+('a'-'A')
	} 
	return c
}
