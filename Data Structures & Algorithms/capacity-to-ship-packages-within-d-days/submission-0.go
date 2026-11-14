func shipWithinDays(weights []int, days int) int {
	minCap,maxCap:=0,0

	for _, weight:=range weights {
		if weight>minCap {
			minCap=weight
		}
		maxCap+=weight
	}

	left,right:=minCap, maxCap

	for left<right {
		mid:=(left+right)/2
		currentLoad:=0
		totalDays:=1

		for _, weight:=range weights {
			if currentLoad+weight>mid {
			totalDays++
			currentLoad=0
		}
		currentLoad+=weight
		}

		if totalDays<=days{
			right=mid
		} else {
			left=mid+1
		}
	}

	return left
}
