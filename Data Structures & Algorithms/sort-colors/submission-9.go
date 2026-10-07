func sortColors(nums []int) {
    first,mid,last:=0,0,len(nums)-1

	for mid<=last {
		switch nums[mid] {
			case 0:
				nums[first],nums[mid]=nums[mid],nums[first]
				first++
				mid++
			case 1:
				mid++
			case 2:
				nums[last],nums[mid]=nums[mid],nums[last]
				last--
		}
	}
}
