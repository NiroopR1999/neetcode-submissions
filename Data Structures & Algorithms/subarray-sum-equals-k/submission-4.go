func subarraySum(nums []int, k int) int {
	prefix:=0
	seen:=make(map[int]int)
	seen[prefix]++
	count:=0
	for _,num:=range nums {
		prefix+=num
		if _,ok:=seen[prefix-k]; ok {
			count+=seen[prefix-k]
		}
		seen[prefix]++

	}
	return count
}
