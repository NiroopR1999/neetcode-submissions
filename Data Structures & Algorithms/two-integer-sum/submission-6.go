func twoSum(nums []int, target int) []int {
    seen:=make(map[int]int)
	res:=[]int{}
	for i,num:=range nums {
		if _,exists:=seen[target-num]; exists {
			 res=append(res,seen[target-num],i)
			 return res
		}
		seen[num]=i
	}
	return res
}
