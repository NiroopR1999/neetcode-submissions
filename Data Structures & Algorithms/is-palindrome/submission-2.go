func isAlphaNumeric(c rune) bool {
	return (c>='a' && c<='z') || (c>='A' && c<='Z') || (c>='0' && c<='9') 
}

func toLower( c rune) byte {
	if(c>='A' && c<='Z') {
		return byte(c + ('a'-'A'))
	}
	return byte(c)
}
func isPalindrome(s string) bool {
	if len(s)==1 {
		return true
	}
	var newStr []byte
	for _,v:=range s {

		if isAlphaNumeric(v) {
			newStr=append(newStr,toLower(v))
		}
	}

	for i,j:=0,len(newStr)-1;i<j;i,j=i+1,j-1 {
		if newStr[i]!=newStr[j] {
			return false
		}
	}

	return true


}
