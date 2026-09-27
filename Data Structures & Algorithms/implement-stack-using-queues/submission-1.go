type MyStack struct {
	 q []int
}

func Constructor() MyStack {
	return MyStack{
		q:[]int{},
	}
}

func (this *MyStack) Push(x int) {
	temp:=[]int{x}
	temp=append(temp,this.q...)
	this.q=temp
}

func (this *MyStack) Pop() int {
	pop:=this.q[0]
	this.q=this.q[1:]
	return pop
}

func (this *MyStack) Top() int {
	return this.q[0]
}

func (this *MyStack) Empty() bool {
	return len(this.q)==0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
