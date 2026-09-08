class Solution {
    /**
     * @param {number} prices
     * @return {number}
     */
    maxProfit(prices) {
       const n=prices.length
       let maxProfit=-Infinity
       let buy=Infinity ,right=0

       while(right<n)
       {
        if(prices[right]<buy)
        {
            buy=prices[right]
        }
        maxProfit=Math.max(maxProfit,prices[right]-buy)
        right++
       }
       return maxProfit
    }
}
