func evalRPN(tokens []string) int {
	res:=[]int{}

	for _, ch :=range tokens{

		if ch !="+" && ch !="-" &&  ch !="*" && ch !="/" {
			val,_:=strconv.Atoi(ch)
			res=append(res,val)
			continue
		}

		right,left:=res[len(res)-1],res[len(res)-2]
		res=res[:len(res)-2]

		switch ch {
			case "+":
				res=append(res,left+right)
			case "-":
				res=append(res,left-right)
			case "*":
				res=append(res,left*right)
			case "/":
				res=append(res,left/right)
		}
	}

	return res[0]
}
