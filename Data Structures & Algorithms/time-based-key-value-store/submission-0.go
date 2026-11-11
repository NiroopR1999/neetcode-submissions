type TimeMap struct {
	store map[string][]pair
}

type pair struct {
	timestamp int
	value string
}
func Constructor() TimeMap {
	return TimeMap{
		store : make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key]=append(this.store[key],
	pair  {
	timestamp :timestamp,
	value : value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result:=""
	history,exists:=this.store[key]; 
	if !exists {
		return result
	}
	left,right:=0,len(history)-1

	for left<=right {

		mid:=(left+right)/2
		if timestamp>=history[mid].timestamp {
			result=history[mid].value
			left=mid+1
		} else {
			right=mid-1
		}
	}
	return result
}
