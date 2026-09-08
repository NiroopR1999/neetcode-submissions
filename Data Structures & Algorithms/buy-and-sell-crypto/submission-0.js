class Solution {
    /**
     * @param {number} prices
     * @return {number}
     */
    maxProfit(prices) {
        let n=prices.length
        let buy=prices[0]
        let maxProfit=0
        for(let i=1;i<n;i++)
        {
            let profit=prices[i]-buy
            maxProfit=Math.max(profit,maxProfit)
            buy=Math.min(prices[i],buy)
        }
        return maxProfit
    }
}
