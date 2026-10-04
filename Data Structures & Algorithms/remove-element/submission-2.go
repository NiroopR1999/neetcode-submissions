func removeElement(nums []int, val int) int {
    writeIndex,readIndex:=0,0

	for readIndex<len(nums) {
		if nums[readIndex]!=val {
			nums[writeIndex]=nums[readIndex]
			writeIndex++
		}
		readIndex++
	}
	return writeIndex
}
