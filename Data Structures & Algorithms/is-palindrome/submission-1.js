class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        let sanitizedString=''

        for(const char of s)
        {
            let e=char.toLowerCase()
            if((e>='a' && e<='z') || (e>='0' && e<='9'))
            {
                sanitizedString+=e
            }
        }

        let left=0
        let right=sanitizedString.length-1
        while(left<right)
        {
            if(sanitizedString[left]!==sanitizedString[right])
            {
                return false
            }
            left++
            right--
        }
        return true

    }
}
