type FreqStack struct {
	freq map[int]int
	freqStack map[int][]int
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack {
		freq : make(map[int]int),
		freqStack : make(map[int][]int),

	}
}

func (this *FreqStack) Push(val int) {
	this.freq[val]++

	f:=this.freq[val]

	this.freqStack[f]=append(this.freqStack[f],val)

	if f>this.maxFreq {
		this.maxFreq=f
	}
}

func (this *FreqStack) Pop() int {
	popStack:=this.freqStack[this.maxFreq]
	popVal:=popStack[len(popStack)-1]
	this.freqStack[this.maxFreq]=this.freqStack[this.maxFreq][:len(this.freqStack[this.maxFreq])-1]
	this.freq[popVal]--
	if len(this.freqStack[this.maxFreq])==0{
		this.maxFreq--
	}
	return popVal
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(val)
 * param2 := obj.Pop()
 */
 