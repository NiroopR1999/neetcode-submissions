class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const map=new Map()
        const n=nums.length
        const res=[]
        const bucket=Array(n+1).fill().map(()=>[])

        for(const num of nums)
        {
            map.set(num,(map.get(num) || 0)+1)
        }
        for(const [num,freq] of map)
        {
            bucket[freq].push(num)
        }
        for(let i=n;i>=0;i--)
        {
            if(bucket[i].length>0)
            {
                for(let j=0;j<bucket[i].length && res.length<k;j++)
                {
                    res.push(bucket[i][j])
                }
            }
        }
        return res
    }
}
