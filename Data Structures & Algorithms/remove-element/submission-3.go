func removeElement(nums []int, val int) int {
    read,write:=0,0
	count:=0
	for read<len(nums) {
		if nums[read]==val {
			read++
		} else {
			// swap
			nums[read],nums[write]=nums[write],nums[read]
			count++
			read++
			write++
		}
	}
	return count
}
