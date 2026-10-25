type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res:=""
	for _, str :=range strs {
		res=res+strconv.Itoa(len(str))+"#"+str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	res:=[]string{}
	i:=0 
	for i<len(encoded) {
		j:=i
		for encoded[j]!='#'{
			j++
		}
		start:=j+1
		end,_:=strconv.Atoi(encoded[i:j])
		res=append(res,encoded[start:start+end])
		i=start+end
	}
	return res
}
