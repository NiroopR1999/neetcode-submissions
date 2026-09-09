func subarraySum(nums []int, k int) int {
	sum,count,i,j:=0,0,0,0

	for {
		
		if j==len(nums) {
			i++
			j=i
			sum=0
		}
		if i==len(nums) {
			break
		}
		sum+=nums[j]
		if sum==k {
			count++
		}
		j++
	}
	return count
}
