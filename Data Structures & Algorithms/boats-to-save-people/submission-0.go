func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left,right,res:=0,len(people)-1,0

	for left<=right {

		if people[left]+people[right]<=limit {
			res++
			left++
			right--
		} else {
			res++
			right--
		}
		
	}
	return res
}
