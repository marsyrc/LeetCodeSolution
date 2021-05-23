package golang_solution.dynamicplanning;

class Solution664 {
    public int strangePrinter(String s) {
        if (s.length() == 0) {
            return 0;
        }
        int n = s.length();
        int[][] dp = new int[n+1][n+1];

        for (int i = 0; i <= n; i++) {
            dp[i][i] = 1;
        }
        for (int l = 2; l <= n; l++) {
            for (int i = 0; i + l - 1 < n; i++) {
                int j = i + l - 1;
                dp[i][j] = 1 + dp[i+1][j];
                for (int k = i + 1; k <= j; k++) {
                    if (s.charAt(k) == s.charAt(i)) {
                        dp[i][j] = Math.min(dp[i][j], dp[i+1][k] + dp[k+1][j]);
                    }
                }
            }
        }
        return dp[0][n-1];
    }
}