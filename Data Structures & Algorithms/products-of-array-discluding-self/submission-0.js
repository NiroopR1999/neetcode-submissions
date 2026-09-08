class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    productExceptSelf(nums) {
       let n=nums.length
       let pre=1
       let post=1
       let res=[]
       for(let i=0;i<n;i++)
       {
         res[i]=pre
         pre*=nums[i]
       }
       for(let i=n-1;i>=0;i--)
       {
         res[i]=res[i]*post
         post*=nums[i]
       }
       return res
}
}
