func trap(height []int) int {
	n:=len(height)
	if n==0 {
		return 0
	}
	
	left:=make([]int,n)
	right:=make([]int,n)
	leftMax:=height[0]
	rightMax:=height[n-1]

	for i:=0;i<n;i++ {
		if height[i]>leftMax{
			leftMax=height[i]
		}
		left[i]=leftMax
	}
	for i:=n-1;i>=0;i--{
		if height[i]>rightMax{
			rightMax=height[i]
		}
		right[i]=rightMax
	}
	total:=0
	for i:=0;i<n;i++ {
		min:=left[i]
		if right[i]<min{
			min=right[i]
		}
		total=total+(min-height[i])
	}
	return total
}
