func removeElement(nums []int, val int) int {
    writeIndex:=0

	for i:=range nums {
		if nums[i]!=val {
			nums[writeIndex]=nums[i]
			writeIndex++
		}
	}
	return writeIndex
}
