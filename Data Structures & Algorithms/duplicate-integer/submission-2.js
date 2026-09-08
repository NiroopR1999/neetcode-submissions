class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {

    //   Method 1 - Time O(n) , Space O(n)
    //   let set= new Set(nums)
    //   return Boolean(nums.length-set.size)

    // Method 2 - Time(nlogn) , Space O(1)
    nums=nums.sort((a,b)=>a-b)
    let n=nums.length
    for(let i=0;i<n-1;i++)
    {
        if(nums[i]===nums[i+1])
        {
            return true
        }
    }
    return false
    }
}
