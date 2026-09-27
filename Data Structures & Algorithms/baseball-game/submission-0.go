func calPoints(operations []string) int {
	res:=[]int{}
	for _, ops :=range operations{
		switch ops {
			case "D": 
					prev:=res[len(res)-1]
					prev=prev*2
					res=append(res,prev)
			case "C":
					res=res[:len(res)-1]
			case "+":
					prev1:=res[len(res)-1]
					prev2:=res[len(res)-2]
					res=append(res,prev1+prev2)
			default:
					val,_:=strconv.Atoi(ops)
					res=append(res,val)

		}
	}
	total:=0
	for i:=range res{
		total+=res[i]
	}
	return total
}
