class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {
        let stack=[]
        let obj={
            '}':'{',
            ")":'(',
            ']':'['
        }
    for(const item of s)
    {
        if(item==='}' || item===')' || item===']')
        {
            if(stack.length===0 || obj[item]!==stack.pop()) return false

        }
        else{
            stack.push(item)
    }}
    return stack.length===0
    }

    }
