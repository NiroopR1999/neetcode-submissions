func majorityElement(nums []int) []int {
	e1,e2:=0,0
	c1,c2:=0,0

	for _,num :=range nums {
		if e1==num && c1>0 {
			c1+=1
		} else if e2==num && c2>0 {
			c2+=1
		} else if c1==0 {
			e1=num
			c1+=1
		} else if c2==0 {
			e2=num
			c2+=1
		} else {
			c1-=1
			c2-=1
		}
	}

	c1,c2=0,0

	for _,num :=range nums {
		if num==e1{
			c1+=1
		}
		if num==e2 {
			c2+=1
		}
	}

	res:=[]int{}

	if c1>len(nums)/3 {
		res=append(res,e1)
	}

	if e1!=e2 && c2>len(nums)/3 {
		res=append(res,e2)
	}
	return res
}
