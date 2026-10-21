func isPalindrome(s string) bool {
	left,right:=0,len(s)-1

	for left<right {
	if !isAlphanumeric(s[left]){
			left++
			continue
		}
	if !isAlphanumeric(s[right]){
			right--
			continue
		}
	if toLower(s[left])!=toLower(s[right]) {
		return false
	}
		left++
		right--
	}
	return true
	
}

func toLower(b byte) byte {
	if( b>='A' && b<='Z') {
		return byte(b+('a'-'A'))
	} 
	return b
}
func isAlphanumeric(b byte) bool {
	return (b>='a' && b<='z') ||
	(b>='A' && b<='Z') || 
	(b>='0' && b<='9') 
}