class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    threeSum(nums) {
        nums = nums.sort((a, b) => a - b)
        let sum
        let res = []
        let n = nums.length
        let i = 0
        while (i < n - 2) {
            if (i > 0 && nums[i] === nums[i - 1]) {
                i++
                continue
            }
            let k = i + 1
            let j = n - 1
            while (k < j) {
                sum = nums[i] + nums[j] + nums[k]
                if (sum === 0) {
                      res.push([nums[i], nums[k], nums[j]]);
                   
                    // Skip duplicates for k
                    while (k < j && nums[k] === nums[k + 1]) k++;
                    // Skip duplicates for j
                    while (k < j && nums[j] === nums[j - 1]) j--;
                       k++;
                    j--;
                  
                } else if (sum > 0) {
                    j--
                } else {
                    k++
                }
            }
            i++
        }
        return res
    }
}