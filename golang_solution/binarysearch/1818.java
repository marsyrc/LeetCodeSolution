package golang_solution.binarysearch;

import java.util.*;

class Solution1818 {
    public int minAbsoluteSumDiff(int[] nums1, int[] nums2) {
        int[] n = new int[nums1.length];
        for (int i = 0; i < n.length; i++) {
            n[i] = nums1[i];
        }
        Arrays.sort(n);
        int max = 0;
        for (int i = 0; i < nums1.length; i++) {
            max = Math.max(max, Math.abs(nums1[i] - nums2[i]) - fun(n, nums2[i]));
        }
        long ans = 0;
        for (int i = 0; i < nums1.length; i++) {
            ans += Math.abs(nums1[i] - nums2[i]);
        }
        ans -= max;
        ans = ans % 1000000007;
        return (int) ans;
    }
    
    int fun(int[] n, int a) {
        int l = 0, r = n.length - 1, mid;
        int la = n[l], ra = n[r];
        while (l <= r) {
            mid = (r + l) / 2;
            if (n[mid] == a) {
                return 0;
            } else if (n[mid] < a) {
                l = mid + 1;
                la = n[mid];
            } else {
                r = mid - 1;
                ra = n[mid];
            }
        }
        return Math.min(Math.abs(a - la), Math.abs(a - ra));
    }
}