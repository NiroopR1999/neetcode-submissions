func largestRectangleArea(heights []int) int {
	stack:=[][]int{}
	maxArea:=0
	for i,height:=range heights {
		start:=i

		for len(stack)>0 && stack[len(stack)-1][1]>height {
			pop:=stack[len(stack)-1]
			stack=stack[:len(stack)-1]
			index:=pop[0]
			value:=pop[1]
			maxArea=max(maxArea,value*(i-index))
			start=index
		}
		stack=append(stack,[]int{start,height})
	}

	for _, pop:=range stack {
		index:=pop[0]
		value:=pop[1]
		maxArea=max(maxArea, value*(len(heights)-index))
	}
	return maxArea
}
