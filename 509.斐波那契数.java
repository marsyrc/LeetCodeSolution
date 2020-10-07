/*
 * @lc app=leetcode.cn id=509 lang=java
 *
 * [509] 斐波那契数
 */

// @lc code=start
class Solution {
    public int fib(int N) {
        if(N == 0)return 0;
        int[] memo  = new int[N+1];
        return helper(memo, N);
    }

    public int helper(int[] memo, int n) {

        if (n == 1 || n == 2) {
            return 1;
        }

        if (memo[n]!= 0 ){
            return memo[n];
        }

        memo[n] = helper(memo, n-1)  + helper( memo, n-2 );
        return memo[n];
    }
}
// @lc code=end
