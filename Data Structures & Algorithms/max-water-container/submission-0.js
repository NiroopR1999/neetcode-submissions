class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights) {

        let n=heights.length
        let i=0
        let j=n-1
        let max=0
        while(i<j)
        {   let min=Math.min(heights[i],heights[j])
            let gap=j-i
            max=Math.max(max,min*gap)
            if(heights[i]<=heights[j])
            {
                i++
            }
            else{
                j--
            }
        }
        return max
    }
}
