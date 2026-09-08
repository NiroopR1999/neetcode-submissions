class Solution {
    /**
     * @param {number[]} height
     * @return {number}
     */
    trap(height) {
        const n=height.length
        let trappedWater=0
        if(n<=2) return trappedWater
        let leftMax=height[0]
        let rightMax=height[n-1]
        let left=1
        let right=n-2
        while(left<=right)
        {
            if(leftMax<rightMax)
            {
                leftMax=Math.max(leftMax,height[left])
                trappedWater+=(leftMax-height[left])
                left++
            }
            else{
                rightMax=Math.max(rightMax,height[right])
                trappedWater+=(rightMax-height[right])
                right--
            }
        }
        return trappedWater
    }
}
