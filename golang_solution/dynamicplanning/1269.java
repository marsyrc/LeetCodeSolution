package golang_solution.dynamicplanning;

class Solution1269 {
    public int numWays(int steps, int arrLen) {
        final int MODULO = 1000000007;
        arrLen = Math.min(arrLen, steps + 1);
        int[]dp = new int[arrLen];
        dp[0] = 1;
        for (int i = 1; i < arrLen; i++) {
            dp[i] = 0;
        }

        for (int i = 1; i <= steps; i++) {
            int[] newdp = new int[arrLen];
            for (int j = 0; j < arrLen; j++) {
                newdp[j] = dp[j];
                if (j > 0) {
                    newdp[j] += dp[j - 1];
                    newdp[j] %= MODULO;
                }
                if (j < arrLen - 1) {
                    newdp[j] += dp[j + 1];
                    newdp[j] %= MODULO;
                }
            }
            dp = newdp;
        }
        return dp[0];
    }
}