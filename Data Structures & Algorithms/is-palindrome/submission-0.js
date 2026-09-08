class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        let filteredStr = '';
        
        // Build the filtered string with alphanumeric characters only
        for (let i = 0; i < s.length; i++) {
            let char = s[i].toLowerCase();
            if ((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
                filteredStr += char;
            }
        }
        
        // Check if the filtered string is a palindrome
        let left = 0;
        let right = filteredStr.length - 1;
        
        while (left < right) {
            if (filteredStr[left] !== filteredStr[right]) {
                return false;
            }
            left++;
            right--;
        }
        
        return true;
    }
}
