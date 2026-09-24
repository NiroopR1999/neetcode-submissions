func maxArea(heights []int) int {
	maxArea:=0
	left,right:=0,len(heights)-1

	for left < right {
		minIndex:=left
		if heights[right] < heights[minIndex] {
			minIndex=right
		}
		area:=heights[minIndex]*(right-left)
		if area>maxArea {
			maxArea=area
		}
		if left==minIndex{
			left++
		} else {
			right--
		}

	}
	return maxArea
}
