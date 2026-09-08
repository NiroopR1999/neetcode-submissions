class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {
        const obj= {']':'[',
        ')':'(',
        '}':"{"}
        const stack=[]
        for(const char of s)
        {
            if(char==='}' || char===']' || char===')')
            {
                if(obj[char]!==stack.pop()) return false
            }
            else
            {
                stack.push(char)
            }
        }
        return stack.length===0
    }
}
