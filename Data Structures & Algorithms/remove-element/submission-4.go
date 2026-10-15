func removeElement(nums []int, val int) int {
    read,write,count:=0,0,0
	for  read<len(nums) {
		if nums[read]==val{
			read++
		} else {
			nums[read],nums[write]=nums[write],nums[read]
			count++
			read++
			write++
		}
	}
	return count
}
