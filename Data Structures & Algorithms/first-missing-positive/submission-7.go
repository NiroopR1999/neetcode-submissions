func firstMissingPositive(nums []int) int {
	index := 0

	for index < len(nums) {
		for nums[index] > 0 &&
			nums[index] <= len(nums) &&
			nums[index] != nums[nums[index]-1] {

			nums[index], nums[nums[index]-1] =
				nums[nums[index]-1], nums[index]
		}
		index++
	}

	for index := 0; index < len(nums); index++ {
		if nums[index] != index+1 {
			return index + 1
		}
	}

	return len(nums) + 1
}