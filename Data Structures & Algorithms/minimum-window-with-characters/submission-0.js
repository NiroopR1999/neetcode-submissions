class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {string}
     */
    minWindow(s, t) {

        let x = s.length
        let y = t.length
        if (y > x || y === 0) return ''

        let charCount = new Map()
        let window = new Map()
        let minLen = Infinity
        let res = [-1, -1];
        let left = 0,
            right = 0
        let met = 0
        for (let i = 0; i < y; i++) {
            charCount.set(t[i], (charCount.get(t[i]) || 0) + 1)
        }

        while (right<x) {
            window.set(s[right], (window.get(s[right]) || 0) + 1)

            if (charCount.has(s[right]) && charCount.get(s[right]) === window.get(s[right])) {
                met++
            }
            while (met === charCount.size) {
                if (right - left + 1 < minLen) {
                    minLen = right - left + 1;
                    res = [left, right];
                }
                // pop from left
                window.set(s[left], (window.get(s[left]) - 1))
                if (charCount.has(s[left]) && window.get(s[left]) < charCount.get(s[left])) met--
                left++
            }
            right++
        }
       return res[0] === -1 ? '' : s.slice(res[0], res[1] + 1);
    }
}