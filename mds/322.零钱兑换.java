import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=322 lang=java
 *
 * [322] 零钱兑换
 */

// @lc code=start
class Solution {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        int max = amount + 1;
        Arrays.fill(dp, max);
        dp[0] = 0;

        for (int i = 0; i < dp.length; ++i) {
            for (int coin : coins) {

                if (i - coin < 0){continue;}

                dp[i] = Math.min(dp[i], 1 + dp[i - coin]);
            }
        }

        return (dp[amount] == max)? -1 : dp[amount];
    }
}
// @lc code=end
