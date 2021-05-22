package dynamicplanning

func longestPalindromeSubseq(s string) int {
	n := len(s)
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, 26)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dp func(i, j, c int) int
	dp = func(i, j, c int) int {
		if memo[i][j][c] != -1 {
			return memo[i][j][c]
		}
		if i >= j {
			memo[i][j][c] = 0
			return 0
		}
		if s[i] != byte('a'+c) {
			memo[i][j][c] = dp(i+1, j, c)
			return memo[i][j][c]
		}
		if s[j] != byte('a'+c) {
			memo[i][j][c] = dp(i, j-1, c)
			return memo[i][j][c]
		}
		res := -1
		for cc := 0; cc < 26; cc++ {
			if cc != c {
				res = max(res, 2+dp(i+1, j-1, cc))
			}
		}
		memo[i][j][c] = res
		return res
	}
	ans := -1
	for cc := 0; cc < 26; cc++ {
		ans = max(ans, dp(0, n-1, cc))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
class Solution {
    public int longestPalindromeSubseq(String s) {
        int n = s.length();
        int[][][] memo = new int[n][n][26];
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                for (int k = 0; k < 26; k++) {
                    memo[i][j][k] = -1;
                }
            }
        }
        int ans = -1;
        for (int c = 0; c < 26; c++) {
            ans = Math.max(ans, dp(0, n - 1, c, memo, s));
        }
        return ans;
    }

    public int dp(int i, int j, int c, int[][][] memo, String s) {
        if (memo[i][j][c] != -1) {
			return memo[i][j][c];
		}
        if (i >= j) {
            memo[i][j][c] = 0;
            return 0;
        }
        if (s.charAt(i) != (char)('a' + c)) {
            memo[i][j][c] = dp(i + 1, j, c, memo, s);
            return memo[i][j][c];
        }
        if (s.charAt(j) != (char)('a' + c)) {
            memo[i][j][c] = dp(i, j - 1, c, memo, s);
            return memo[i][j][c];
        }
        int res = -1;
        for (int delta = 0; delta < 26; delta++) {
            if (c != delta) {
                res = Math.max(res, 2 + dp(i + 1, j - 1, delta, memo, s));
            }
        }
        memo[i][j][c] = res;
        return res;
    }
}
*/
