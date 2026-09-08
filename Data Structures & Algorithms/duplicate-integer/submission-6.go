func hasDuplicate(nums []int) bool {
    res:=make(map[int]int)

	for _,v:= range nums {
		if _,ok := res[v]; ok {
			return true
		}
		res[v]=res[v]+1
	}
	return false
}
