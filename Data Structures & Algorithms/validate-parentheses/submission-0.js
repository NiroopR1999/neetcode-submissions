class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {

        let res=[]
         const obj = {
            ')': '(',
            ']': '[',
            '}': '{'
        };

        for(const item of s)
        {
           if(item===']' || item===')' || item==='}')
           {
            if(res.length===0) return false
            if(obj[item]!==res.pop()) return false
           }
           else{
            res.push(item)
           }
        }
       return  res.length===0
    }
}
