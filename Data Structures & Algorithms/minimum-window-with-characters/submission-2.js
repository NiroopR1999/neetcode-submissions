class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {string}
     */
    minWindow(s, t) {
        // Length of the input strings
        let x = s.length;
        let y = t.length;
        // If t is longer than s or t is empty, return an empty string
        if (y > x || y === 0) return '';

        // Maps to keep track of character counts
        let charCount = new Map();  // Required characters and their counts in t
        let window = new Map();     // Characters in the current window in s
        // Initialize variables to track the minimum window length and its bounds
        let minLen = Infinity;
        let res = [-1];
        // Pointers for the sliding window
        let left = 0, right = 0;
        // Counter for how many characters in t are fully matched in the current window
        let met = 0;

        // Fill charCount with the count of each character in t
        for (let i = 0; i < y; i++) {
            charCount.set(t[i], (charCount.get(t[i]) || 0) + 1);
        }

        // Expand the right end of the window
        while (right < x) {
            // Add current character to the window map
            window.set(s[right], (window.get(s[right]) || 0) + 1);

            // If the current character is in t and its count in the window matches the required count, increment met
            if (charCount.has(s[right]) && charCount.get(s[right]) === window.get(s[right])) {
                met++;
            }

            // Contract the left end of the window as long as all characters in t are fully matched
            while (met === charCount.size) {
                // Update the minimum length and result bounds if the current window is smaller
                if (right - left + 1 < minLen) {
                    minLen = right - left + 1;
                    res = [left, right];
                }
                // Remove the character at the left end of the window
                window.set(s[left], window.get(s[left]) - 1);
                // If this character is in t and its count in the window is now less than required, decrement met
                if (charCount.has(s[left]) && window.get(s[left]) < charCount.get(s[left])) {
                    met--;
                }
                // Move the left pointer to the right
                left++;
            }
            // Move the right pointer to the right
            right++;
        }

        // If no valid window was found, return an empty string, otherwise return the minimum window
        return res[0] === -1 ? '' : s.slice(res[0], res[1] + 1);
    }
}
