class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    lengthOfLongestSubstring(s) {
        const n=s.length
        const map= new Map()
        let i=0, j=0 ,max=0

        while(j<n)
        {
            if(map.has(s[j]))
            {
                map.delete(s[i])
                i++
            }
            else{
                map.set(s[j],1)
                max=Math.max(j-i+1,max)
                j++
            }
            
        }
        return max
    }
}
