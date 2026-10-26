type LRUCache struct {
    capacity int
	cache map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	key int
	value int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	head:=&Node{
	}
	tail:=&Node{

	}
	head.Next=tail
	tail.Prev=head
    return LRUCache{
		capacity:capacity,
		cache:make(map[int]*Node),
		head:head,
		tail:tail,
	}
}
// remove
func (l *LRUCache) remove(node *Node){
	node.Prev.Next=node.Next
	node.Next.Prev=node.Prev
}
// movetoFront
func (l *LRUCache) moveToFront(node *Node){
	l.remove(node)
	l.addFront(node)
}
// addFront
func (l *LRUCache) addFront(node *Node){
	node.Next=l.head.Next
	node.Prev=l.head
	l.head.Next.Prev=node
	l.head.Next=node
}

func (this *LRUCache) Get(key int) int {
    node,exists:=this.cache[key]
	if !exists{
		return -1
	}
	this.moveToFront(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// case : key already exists
	node,exists:=this.cache[key]
	if exists{
		node.value=value
		this.moveToFront(node)
		return
	}

	// case: new key
	if len(this.cache)>=this.capacity {
		lru:=this.tail.Prev
		this.remove(lru)
		delete(this.cache,lru.key)
	}

	node=&Node{
		value:value,
		key:key,
	}
	this.addFront(node)
	this.cache[key]=node

}
