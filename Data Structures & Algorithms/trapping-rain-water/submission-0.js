class Solution {
    /**
     * @param {number[]} height
     * @return {number}
     */
    trap(height) {
        let n=height.length
        let leftMax=height[0]
        let rightMax=height[n-1]
        let trappedWater=0
        let i=1
        let j=n-2

        while(i<=j)
        {
            if(leftMax<rightMax)
            {
                if(height[i]>leftMax) leftMax=height[i]
                 trappedWater+=leftMax-height[i]
                 i++
            }
            else{
                if(height[j]>rightMax) rightMax=height[j]
                 trappedWater+=rightMax-height[j]
                 j--
            }
           
        }
        return trappedWater
    }
}
