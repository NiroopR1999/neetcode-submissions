func maxArea(heights []int) int {
	left,most,right:=0,0,len(heights)-1

	for left<right {
		if heights[left]<heights[right] {
			most=max(most,heights[left]*(right-left))
			left++
		} else {
			most=max(most,heights[right]*(right-left))
			right--
		}
	}

	return most
}
