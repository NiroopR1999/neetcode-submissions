class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
        const set= new Set(nums)
        let max=0
        let count=0
        for(const num of set)
        {   count=0
            if(!set.has(num-1))
            {
                count++
                let current=num
                while(set.has(current+1))
                {
                    current++
                    count++
                }
                max=Math.max(max,count)
            }
        }
        return max
    }
}
