func twoSum(nums []int, target int) []int {

    res:=make(map[int]int)

	for i, v:= range nums {
		rem:=target-v
		if _,ok:=res[rem]; ok {
			return []int{res[rem],i}
		} else {
			res[v]=i
		}
	}
	return nil
}
 