func minWindow(s string, t string) string {
	left,right,start,end:=0,0,0,0
	freq:=[128]int{}
	need:=len(t)
	for i:=range t {
		freq[t[i]-'A']++
	}
	
	for right<len(s) {
		if freq[s[right]-'A']>0 {
			need--
		}
		freq[s[right]-'A']--
		for need==0 {
			if end==0|| right-left+1<end-start {
				start=left
				end=right+1
			}
			freq[s[left]-'A']++
			if freq[s[left]-'A']>0 {
				need++
			}
			left++
		}
		right++
	}
	return s[start:end]
}
