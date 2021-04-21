package golang_solution.binarysearch;

class Solution1712 {
    public int waysToSplit(int[] nums) {
        int n = nums.length;
        int[] prefix = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            prefix[i] = prefix[i - 1] + nums[i - 1];
        }
        int res = 0;
        for (int i = 1; i < n - 1; i++) {
            if (prefix[i] * 3 > prefix[n]) {
                break;
            }
            // find right boundary
            int l = i + 1, r = n - 1;
            while (l < r) {
                int mid = (l + r + 1) / 2;
                int b = prefix[mid] - prefix[i];
                int c = prefix[n] - prefix[mid];
                if (b <= c) {
                    l = mid;
                } else {
                    r = mid - 1;
                }
            }
            // find left boundary
            int ll = i + 1;
            int rr = n - 1;
            while (ll < rr) {
                int mid = (ll + rr) / 2;
                int a = prefix[i];
                int b = prefix[mid] - prefix[i];
                if (b >= a) {
                    rr = mid;
                } else {
                    ll = mid + 1;
                }
            }
            res += r - ll + 1;
            res %= 1e9 + 7;
        }
        return res;
    }
}