class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        let n = nums.length;
        let map = new Map();
        let bucket = Array(n + 1).fill().map(() => []);
        let res = [];

        for(let i = 0; i < n; i++) {
            map.set(nums[i], (map.get(nums[i]) || 0) + 1);
        }

        for(let [num, freq] of map) {
            bucket[freq].push(num);
        }
            console.log(bucket)
        for(let i = bucket.length - 1; i >= 0; i--) {
            if(bucket[i].length > 0) {
                for(let j = 0; j < bucket[i].length && res.length < k; j++) {
                    res.push(bucket[i][j]);
                }
            }
        }

        return res;
    }
}
