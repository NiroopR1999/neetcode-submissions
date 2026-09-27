func isValid(s string) bool {
    res:=[]rune{}
	match:=map[rune]rune{
		')':'(',
		']':'[',
		'}':'{',
	}
	for _,ch:=range s {

		if ch==')' || ch=='}' || ch==']' {
				if len(res)==0 {
					return false
				}
				if res[len(res)-1]!=match[ch]  {
					return false
				}
				res=res[:len(res)-1]
				continue
		}

		res=append(res,ch)

	}
	return len(res)==0
}
