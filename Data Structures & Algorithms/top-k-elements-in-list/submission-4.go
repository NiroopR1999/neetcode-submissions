func topKFrequent(nums []int, k int) []int {
	if len(nums)==0 {
		return []int{}
	}
	freq:=make(map[int]int)

	// create frequency of each number
	for _,num:=range nums {
		freq[num]++
	}
	
	// add the numbers to frequency bucket
	bucket:=make([][]int,len(nums)+1)

	for key,value:=range freq {
		bucket[value]=append(bucket[value],key)
	}

	res:=[]int{}
	for i:=len(bucket)-1;i>=0;i-- {
		
		for _, value:=range bucket[i] {
			res=append(res,value)
			if len(res)==k {
				return res
			}
		}
	}
	return res
}
