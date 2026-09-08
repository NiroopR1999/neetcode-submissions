class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    threeSum(nums) {
        const n=nums.length
        const res=[]
        nums=nums.sort((a,b)=>a-b)
        let i=0
        while(i<n-2)
        {
            if(i>0 && nums[i]===nums[i-1])
            {   
                i++
                continue
            }

            let j=i+1
            let k=n-1
            while(j<k)
            {
                const sum=nums[i]+nums[j]+nums[k]

                if(sum===0)
                {
                    res.push([nums[i],nums[j],nums[k]])

                    while(j<k && nums[j]===nums[j+1]) j++
                    while (j<k && nums[k]===nums[k-1]) k--
                    j++
                    k--
                }
                else if(sum>0)
                {
                    k--
                }
                else{
                    j++
                }
            }
            i++

        }
        return res
    }
}
