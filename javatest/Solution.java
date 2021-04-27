package javatest;

import java.util.*;


class Solution {
    public int maxFrequency(int[] nums, int k) {
        Arrays.sort(nums);
        long total = 0;
        int res = 0;
        int l = 0;
        for (int r = 1; r < nums.length; r++) {
            total += (nums[r] - nums[r - 1]) * (r - l);
            while (total > k) {
                total -= nums[r] - nums[l];
                l++;
            }

            res = Math.max(res, r - l + 1);
        }
        return res;
    }
}
