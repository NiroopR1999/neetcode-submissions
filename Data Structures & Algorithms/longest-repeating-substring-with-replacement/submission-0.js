class Solution {
    /**
     * @param {string} s
     * @param {number} k
     * @return {number}
     */
    characterReplacement(s, k) {
        let maxCount=0
        let maxLength=0
        let map = new Map()
        let i=0
        let j=0
        let n=s.length
        while(j<n)
        {

            map.set(s[j], (map.get(s[j]) || 0) + 1);
            maxCount = Math.max(maxCount, map.get(s[j]));

            if(j-i+1-maxCount>k)
            {
                map.set(s[i], map.get(s[i]) - 1);
                 i++
            }

            maxLength=Math.max(maxLength,j-i+1)
            j++
        }
        return maxLength
    }
}
