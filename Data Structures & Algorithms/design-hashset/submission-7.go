type MyHashSet struct {
	data [1_000_001]bool
}

func Constructor() MyHashSet {
    return MyHashSet{
		data: [1_000_001]bool{},
	}
}

func (this *MyHashSet) Add(key int) {
    this.data[key]=true
}

func (this *MyHashSet) Remove(key int) {
    this.data[key]=false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.data[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 