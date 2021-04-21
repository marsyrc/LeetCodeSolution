package golang_solution.prefixsum;

class Solution1314 {
	public int[][] matrixBlockSum(int[][] mat, int K) {
		int m = mat.length, n = mat[0].length;
		int[][] sum = new int[m + 1][n + 1];
		int[][] res = new int[m][n];
		for (int i = 1; i <= m; i++) {
			for (int j = 1; j <= n; j++) {
				sum[i][j] = sum[i - 1][j] + sum[i][j - 1] - sum[i - 1][j - 1] + mat[i - 1][j - 1];
			}
		}
		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				int ux = Math.min(i + K + 1, m);
				int uy = Math.min(j + K + 1, n);
				int dx = Math.max(i - K, 0);
				int dy = Math.max(j - K, 0);
				res[i][j] = sum[ux][uy] - sum[ux][dy] - sum[dx][uy] + sum[dx][dy];
			}
		}
		return res;
	}
}