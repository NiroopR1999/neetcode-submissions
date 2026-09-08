class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    threeSum(nums) {
        nums.sort((a, b) => a - b); // Sort the array
        const triplets = [];

        for (let i = 0; i < nums.length - 2; i++) {
            // Skip duplicates for i
            if (i > 0 && nums[i] === nums[i - 1]) continue;

            let j = i + 1; // Left pointer
            let k = nums.length - 1; // Right pointer

            while (j < k) {
                const sum = nums[i] + nums[j] + nums[k];

                if (sum === 0) {
                    triplets.push([nums[i], nums[j], nums[k]]);
     // Move both pointers after processing
                    j++;
                    k--;
                    // Skip duplicates for j using previous element
                    while (j < k && nums[j] === nums[j - 1]) j++;
                    // Skip duplicates for k using next element
                    while (j < k && nums[k] === nums[k + 1]) k--;

                   
                } else if (sum < 0) {
                    j++; // We need a larger sum, move left pointer right
                } else {
                    k--; // We need a smaller sum, move right pointer left
                }
            }
        }
        return triplets;
    }
}
