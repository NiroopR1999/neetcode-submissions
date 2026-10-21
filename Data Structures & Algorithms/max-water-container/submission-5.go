func maxArea(heights []int) int {
	most:=0
	left,right:=0,len(heights)-1

	for left<right {
		most=max(most,min(heights[left],heights[right])*(right-left))
		if heights[left]<heights[right] {
			left++
		} else {
			right--
		}
	}

	return most
}
