class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {

    // Method 1 - Time O(n) , Time (1)

     if(s.length !== t.length) return false;

        let s1 = s.toLowerCase();
        let t1 = t.toLowerCase();

        let arr = Array(26).fill(0); 

        for(let i = 0; i < s1.length; i++) {
            arr[s1[i].charCodeAt() - 97]++; 
        }

        for(let i = 0; i < t1.length; i++) {
            arr[t1[i].charCodeAt() - 97]--; 
        }

        return arr.every(val => val === 0);
    }
}
