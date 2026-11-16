func maxArea(heights []int) int {
	maxWater:=0
	left,right:=0,len(heights)-1

	for left<right {
		if heights[left]<heights[right] {
			maxWater=max(maxWater, heights[left]*(right-left))
			left++
		} else {
			maxWater=max(maxWater, heights[right]*(right-left))
			right--
		}
	}

	return maxWater
}
