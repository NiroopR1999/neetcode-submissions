class Solution {
    /**
     * @param {string} s
     * @param {number} k
     * @return {number}
     */
    characterReplacement(s, k) {
       const n=s.length
       let i=0, j=0, map = new Map()
       let max=-Infinity
        let maxF=-Infinity

       while(j<n)
       {
          map.set(s[j], (map.get(s[j]) || 0) + 1);
          maxF=Math.max(maxF,map.get(s[j]))

        if (j - i + 1 - maxF <= k) 
        {
            max=Math.max(max,j-i+1)
        }
        else{
            map.set(s[i],map.get(s[i])-1)
            i++
        }
        j++
       }
        return max
    }
   
}
