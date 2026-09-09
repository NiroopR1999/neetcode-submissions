func subarraySum(nums []int, k int) int {
	freq:=map[int]int{0:1}
	prefixSum:=0
	count:=0
	for _,v:=range nums{
		prefixSum+=v
		count+=freq[prefixSum-k]
		freq[prefixSum]+=1
	}
	return count
}
