func sortArray(nums []int) []int {
    if len(nums)<= 1 {
		return nums
	}
	left:=sortArray(nums[:len(nums)/2])
	right:=sortArray(nums[len(nums)/2:])

	return merge(left,right)
}

func merge(left, right []int) []int {
	res:=make([]int,0,len(left)+len(right))
	i,j:=0,0

	for i<len(left) && j <len(right) {

		if left[i]<right[j] {
			res=append(res,left[i])
			i++
		} else {
			res=append(res,right[j])
			j++
		}
	}
	res=append(res,left[i:]...)
	res=append(res,right[j:]...)
	return res
}