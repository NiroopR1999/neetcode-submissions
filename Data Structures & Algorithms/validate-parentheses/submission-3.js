class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {

        let obj={
            '}':'{',
            ']':'[',
            ")":'('
        }
        const stack=[]

        for(const item of s)
        {
            if(item===']' || item==='}' || item===')' )
            {
                if(obj[item]!==stack.pop()) return false
            }
            else{
                stack.push(item)
            }
        }
        return stack.length===0
    }
}
