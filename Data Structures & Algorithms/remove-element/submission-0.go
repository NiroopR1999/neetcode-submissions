func removeElement(nums []int, val int) int {
	writeIndex:=0

	for i, v := range nums {
		if v==val {
			continue
		} else {
			nums[writeIndex]=nums[i]
			writeIndex+=1
		}
 	}
	return len(nums[:writeIndex])
}
