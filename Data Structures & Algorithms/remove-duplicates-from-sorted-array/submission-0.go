func removeDuplicates(nums []int) int {
	write,read:=1,1

	for read<len(nums) {
		if nums[read]!=nums[read-1] {
			nums[write]=nums[read]
			write++
		}
		read++
	}
	return write
}
