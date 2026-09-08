class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {

    // Method 1 - Time O(n) , Time (1)

    if(s.length!==t.length) return false
    
    const s1=s.toLowerCase()
    const t1=t.toLowerCase()
    const arr=Array(26).fill(0)

    for(const char of s1)
    {
        arr[char.charCodeAt()-97]++
    }
    for(const char of t1)
    {
        arr[char.charCodeAt()-97]--
    }

    return arr.every(val=>val===0)
    }
}
