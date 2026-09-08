class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {

        // Method 1 - Time O(nlogn) , Space (1)
         // Convert strings to arrays of characters and sort them
    let s1 = s.split('').sort().join('');
    let t1 = t.split('').sort().join('');

        for(let i=0;i<Math.max(s.length,t.length);i++)
        {
            if(t1[i]!==s1[i])
            {
                return false
            }
        }
        return true
    }
}
