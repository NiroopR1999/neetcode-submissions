class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */
    checkInclusion(s1, s2) {
        if(s1.length>s2.length) return false

        let arr = Array(26).fill(0)

        for (let i = 0; i < s1.length; i++) {
            arr[s1.charCodeAt(i) - 'a'.charCodeAt(0)]++;
            arr[s2.charCodeAt(i) - 'a'.charCodeAt(0)]--;
        }

        if (arr.every(item => item === 0)) return true

        for (let i = s1.length; i < s2.length; i++) {
            arr[s2[i].charCodeAt() - 97]--
            arr[s2.charCodeAt(i - s1.length) - 97]++
            if (arr.every(item => item === 0)) return true
        }
        return false
    }
}