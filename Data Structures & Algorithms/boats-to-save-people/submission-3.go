func numRescueBoats(people []int, limit int) int {
	count:=0
	left,right:=0,len(people)-1 
	sort.Ints(people)
	for left<=right {
		if people[left]+people[right]<=limit {
			left++
			right--
		} else if people[left]>people[right] {
			left++
		} else {
			right--
		}
			count++
	}
	return count
}
