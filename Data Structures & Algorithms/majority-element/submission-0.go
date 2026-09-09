func majorityElement(nums []int) int {
    var res= map[int]int {}

	for _,v:=range nums{
		res[v]++
	}

	for k,v:=range res{
		if v>len(nums)/2 {
			return k
		}
	}
	return 0
}
