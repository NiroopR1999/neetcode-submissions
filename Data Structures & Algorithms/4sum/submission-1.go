func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res:=[][]int{}

	for i:=0;i<len(nums)-3;i++{
		if i>0 && nums[i]==nums[i-1]{
			continue
		}

		for j:=i+1;j<len(nums)-2;j++{
			if j>i+1 &&  nums[j]==nums[j-1]{
				continue
			}

			left,right:=j+1,len(nums)-1

			for left<right{
				sum:=nums[i]+nums[j]+nums[left]+nums[right]
				
				if sum==target {
					res=append(res,[]int{nums[i],nums[j],nums[left],nums[right]})
					left+=1
					right-=1

					for left<right && nums[left]==nums[left-1] {
						left+=1
					}

					for left<right && nums[right]==nums[right+1]{
						right-=1
					}
				} else if sum>target {
					right-=1
				} else {
					left+=1
				}
			}
		}
	}
	return res
}
