class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    productExceptSelf(nums) {
        const n=nums.length
        const res=[]
        let prefix=1
        let postfix=1

        for(let i=0;i<n;i++)
        {
            res[i]=prefix
            prefix*=nums[i]
        }

        for(let i=n-1;i>=0;i--)
        {
            res[i]=res[i]*postfix
            postfix*=nums[i]
        }
        return res
    }
}
