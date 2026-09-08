class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let n=nums.length
        let map= new Map()
        for(let i=0;i<n;i++)
        {
           if(map.has(nums[i]))
           {
            return [map.get(nums[i]),i]
           }
           map.set(target-nums[i],i)
        }
    }
}
