type MinStack struct {
	s []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		s :[]int{},
		min:[]int{},
	}
}

func (this *MinStack) Push(val int) {
	this.s=append(this.s,val)
	if len(this.min)==0{
		this.min=append(this.min,val)
	}else {
		minVal:=this.min[len(this.min)-1]
		if val<minVal {
			minVal=val
		}
		this.min=append(this.min,minVal)
	}
}

func (this *MinStack) Pop() {
	this.s=this.s[:len(this.s)-1]
	this.min=this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
 	return this.min[len(this.min)-1]
}
