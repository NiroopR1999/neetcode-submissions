func sortColors(nums []int) {
    left,mid,right:=0,0,len(nums)-1

	for mid<=right { 
		switch  nums[mid] {
			case 0 :
				nums[mid],nums[left]=nums[left],nums[mid]
				left+=1
				mid+=1
			case 1:
				mid+=1
			case 2:
				nums[mid],nums[right]=nums[right],nums[mid]
				right-=1
		}
	}
}
