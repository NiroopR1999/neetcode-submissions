func subarraySum(nums []int, k int) int {
	seen:=make(map[int]int)
	prefix:=0
	seen[prefix]++
	count:=0
	for _,num:=range nums {
		prefix+=num
		if val,ok:=seen[prefix-k];ok {
			count+=val
		}
		seen[prefix]++
	}
	return count
}
