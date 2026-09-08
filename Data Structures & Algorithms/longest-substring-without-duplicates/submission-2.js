class Solution {
    /**
     * @param {string} s
     * @return {number}
     */
    lengthOfLongestSubstring(s) {
        // Method 1

        let map = new Map();
        let count = 0;
        let i = 0, j = 0;
        while (j < s.length) {
            if (map.has(s[j])) {
                map.delete(s[i]);
                i++;
            } else {
                map.set(s[j], true);
                count = Math.max(count, j - i + 1);
                j++;
            }
        }

        return count;
    }
}
