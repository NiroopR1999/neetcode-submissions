type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{
	stack:[]int{},
}
}

func (this *MyQueue) Push(x int) {
	this.stack=append(this.stack,x)
}

func (this *MyQueue) Pop() int {
	pop:=this.stack[0]
	this.stack=this.stack[1:]
	return pop
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack)==0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
