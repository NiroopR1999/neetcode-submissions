class Solution {
    /**
     * @param {number} n
     * @return {string[]}
     */
    generateParenthesis(n) {
        const res=[]

        const generate=(s='',open=0,close=0)=>
        {
            if(open===n && close===n) return res.push(s)

            if(open<n)
            {
                generate(s+'(',open+1,close)
            }
            if(close<open)
            {
                generate(s+')',open,close+1)
            }
        }
        generate()
        return res
    }
}
