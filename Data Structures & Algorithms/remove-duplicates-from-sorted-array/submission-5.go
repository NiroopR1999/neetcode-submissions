func removeDuplicates(nums []int) int {
	read,write:=0,0

	for read<len(nums) {
		if nums[read]!=nums[write]{
			write+=1
			nums[write]=nums[read]
		}
		read++
	}
	return write+1
}
