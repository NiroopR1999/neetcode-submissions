class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    largestRectangleArea(heights) {
        const n=heights.length
        const stack=[]
        let maxArea=0

        for(let i=0;i<n;i++)
        {   
            let startIndex=i
            while(stack.length>0 && heights[i]<stack[stack.length-1][1])
            {
                let top=stack.pop()
                maxArea=Math.max(maxArea,top[1]*(i-top[0]))
                startIndex=top[0]
            }
            stack.push([startIndex,heights[i]])
        }
        for(const item of stack)
        {
           maxArea=Math.max(maxArea,item[1]*(n-item[0]))
        }
        return maxArea
    }
}
