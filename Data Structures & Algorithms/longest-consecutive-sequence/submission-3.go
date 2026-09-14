func longestConsecutive(nums []int) int {
	longest:=0
	set:=make(map[int]bool)

	for _,v:=range nums {
		set[v]=true
	}

	for _,v:=range nums{

		if !set[v-1] {
			count:=1

			for set[v+count]{
				count++
			}

			if count>longest {
				longest=count
			}
		}
	}

	return longest
}
