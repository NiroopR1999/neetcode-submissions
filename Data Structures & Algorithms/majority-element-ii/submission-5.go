func majorityElement(nums []int) []int {
	cd1,cd2,c1,c2:=0,0,0,0
	n:=len(nums)
	for _,num:=range nums{
		if num==cd1 && c1>0{
			c1++
		} else if num==cd2 && c2>0 {
			c2++
		} else if c1==0 {
			cd1=num
			c1++
		} else if c2==0 {
			cd2=num
			c2++
		} else {
			c1--
			c2--
		}
	}	

	c1,c2=0,0

	for _,num:=range nums {
		if num==cd1{
			c1++
		}
		if num==cd2 {
			c2++
		}
		}
		res:=[]int{}

		if c1>n/3 {
			res=append(res,cd1)
		}
		if c2>n/3 {
			res=append(res,cd2)
		}
		return res

}
