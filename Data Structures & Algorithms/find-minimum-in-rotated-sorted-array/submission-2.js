class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findMin(nums) {
        let n=nums.length
        let left=0, right=n-1
        let min=Infinity

        while(left<=right)
        {
            let mid=Math.floor((left+right)/2)

            if(nums[left]<=nums[mid])
            {
                min=Math.min(min,nums[left])
                left=mid+1
            }
           if(nums[mid]<=nums[right])
            {
                min=Math.min(min,nums[mid])
                right=mid-1
            }
        }
        return min
    }
}
