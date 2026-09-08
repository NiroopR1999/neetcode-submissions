class Solution {
    /**
     * @param {number[]} temperatures
     * @return {number[]}
     */
    dailyTemperatures(temperatures) {
        let n=temperatures.length
        let res=new Array(n).fill(0)
        let stack=[]

        for(let i=0;i<n;i++)
        {
            while(stack.length>0 && temperatures[i]>temperatures[stack[stack.length-1]])
            {
             let index=stack.pop()
             res[index]=i-index
            }
            stack.push(i)
        }

        return res

    }
}
