func topKFrequent(nums []int, k int) []int {
	freq:=make(map[int]int)

	// find frequency of each numbers
	for _,num :=range nums {
		freq[num]++
	}

	// create buckets for each freq 

	buckets:=make([][]int,len(nums)+1)

	for key,value :=range freq {
		buckets[value]=append(buckets[value],key)
	}

	// iterate from end of array
	res:=[]int{}
	for i:=len(buckets)-1;i>=0;i-- {
		
		for _, val :=range buckets[i] {
			res=append(res,val)
			if len(res)==k {
				return res
			}
		}
	}
	return res
}
