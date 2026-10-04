func majorityElement(nums []int) []int {
	candidate1,candidate2,count1,count2:=0,0,0,0

	for _,v:=range nums {
		if v==candidate1 && count1>0 {
			count1++
		} else if v==candidate2 && count2>0 {
			count2++
		} else if  count1==0 {
			candidate1=v
			count1++
		} else if  count2==0 {
			candidate2=v
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1,count2=0,0
	res:=make([]int,0,2)


	for _,v:=range nums {
		if v==candidate1{
			count1++
		} 
		if v==candidate2{
			count2++
		} 
	}

	if count1>len(nums)/3 {
		res=append(res,candidate1)
	}
	if count2>len(nums)/3 {
		res=append(res,candidate2)
	}
	return res
}
