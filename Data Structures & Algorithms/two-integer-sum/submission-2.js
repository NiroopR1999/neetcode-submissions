class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {

        const map= new Map()
        const n=nums.length
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
