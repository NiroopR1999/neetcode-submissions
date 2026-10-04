func subarraySum(nums []int, k int) int {
	prefixSum:=make(map[int]int)
	prefix,count:=0,0
	prefixSum[prefix]++
	for _,num:=range nums {
		prefix+=num
		count+=prefixSum[prefix-k]
		prefixSum[prefix]++
	}
	return count
}
