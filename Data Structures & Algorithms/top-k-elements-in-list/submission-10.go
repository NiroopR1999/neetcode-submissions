func topKFrequent(nums []int, k int) []int {
	freq:=make(map[int]int)

	for _,num :=range nums {
		freq[num]++
	}

	// create bucket
	bucket:=make([][]int,len(nums)+1)

	for num, index:=range freq {
		bucket[index]=append(bucket[index],num)
	}

	res:=[]int{}
	for i:=len(bucket)-1;i>=0;i--{
		for _,num:=range bucket[i] {
			res=append(res,num)
			if len(res)==k {
				return res
			}
		}
	}
	return res
}
