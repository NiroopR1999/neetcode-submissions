class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
      let set= new Set(nums)

      return Boolean(nums.length-set.size)
    }
}
