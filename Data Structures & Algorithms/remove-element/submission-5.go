func removeElement(nums []int, val int) int {
    
	read,write,count:=0,0,0

	for read<len(nums) {
		if nums[read]!=val{
			nums[write]=nums[read]
			count++
			write++
		}
		read++
	}
	return count
}
