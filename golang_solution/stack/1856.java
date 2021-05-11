import java.util.*;

class Solution {
    final long mod = (long) (1e9 + 7);

    public int maxSumMinProduct(int[] nums) {
        int n = nums.length;
        long ans = 0;

        int[] left = new int[n];
        int[] right = new int[n];
        Arrays.fill(right, n - 1);

        Stack<Integer> s = new Stack<>();
        for (int i = 0; i < n; ++i) {
            while (!s.isEmpty() && nums[s.peek()] > nums[i]) {
                right[s.peek()] = i - 1;
                s.pop();
            }

            if (!s.isEmpty()) {
                left[i] = s.peek() + 1;
            }
            s.push(i);
        }

        long[] prefix = new long[n + 1];
        for (int i = 1; i <= n; i++) {
            prefix[i] = prefix[i - 1] + nums[i - 1];
        }

        for (int i = 0; i < n; i++) {
            ans = Math.max(ans, (prefix[right[i] + 1] - prefix[left[i]]) * nums[i]);
        }
        return (int) (ans % mod);
    }
}