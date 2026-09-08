class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
        let set= new Set(nums)
        let max=0, count=0 , n=nums.length
        for(let i=0;i<n;i++)
        {   
            let num=nums[i]
            if(!set.has(num-1))
            {
                count++
                while(set.has(num+1))
                {   
                    num++
                    count++
                }
                max=Math.max(max,count)
            }
            count=0

        }
        return max
    }
}
