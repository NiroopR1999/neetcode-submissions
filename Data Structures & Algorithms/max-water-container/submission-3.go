func maxArea(heights []int) int {
	most:=0
	left,right:=0,len(heights)-1

	for left<right {
		area:=min(heights[left],heights[right])*(right-left)
		most=max(most,area)
		if heights[left]>heights[right] {
			right--
		} else {
			left++
		}
	}
	return most
}
