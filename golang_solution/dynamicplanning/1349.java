package golang_solution.dynamicplanning;

class Solution {
    public int maxStudents(char[][] seats) {
        int m = seats.length;
        int n = seats[0].length;

        int[] valid = new int[m];
        int stateSize = (1 << n);
        int[][] dp = new int[m][stateSize];
        int ans = 0;

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                valid[i] = (valid[i] << 1) + (seats[i][j] == '.' ? 1 : 0);
            }
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < stateSize; j++) {
                dp[i][j] = -1;
            }
        }

        for (int i = 0; i < m; i++) {
            for (int j = 0; j < stateSize; j++) {
                if (((j & valid[i]) == j) && ((j & (j >> 1)) == 0)) {
                    if (i == 0) {
                        dp[i][j] = Integer.bitCount(j);
                    } else {
                        for (int k = 0; k < stateSize; k++) {
                            if ((j & (k >> 1)) == 0 && ((j >> 1) & k) == 0 && (dp[i - 1][k] != -1)) {
                                dp[i][j] = Math.max(dp[i - 1][k] + Integer.bitCount(j), dp[i][j]);
                            }
                        }
                    }
                    ans = Math.max(dp[i][j], ans);
                }
            }
        }
        return ans;
    }

    public static void main(String[] args) {

    }
}
