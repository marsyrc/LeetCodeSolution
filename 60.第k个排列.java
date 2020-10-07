/*
 * @lc app=leetcode.cn id=60 lang=java
 *
 * [60] 第k个排列
 */

// @lc code=start
class Solution {
    private int[] factorial; //阶乘数组

    private boolean[] used;

    private int n;
    private int k;

    private void calculateFactorial(int n) {
        factorial = new int[n + 1];
        factorial[0] = 1;
        for (int i = 1; i <= n; i++) {
            factorial[i] = factorial[i - 1] * i;
        }
    }

    public String getPermutation(int n, int k) {
        this.n = n;
        this.k = k;
        calculateFactorial(n);

        used = new boolean[n + 1];
        Arrays.fill(used,false);

        StringBuilder path = new StringBuilder();
        dfs(0 , path);

        return  path.toString();
    }

    private void dfs(int index, StringBuilder path){
        if (index == n){
            return;
        }
        int count = factorial[n - 1 - index];

        for (int i = 1; i <= n; i++) {
            if (used[i]){
                continue;
            }

            if (count < k){
                k -= count;
                continue;
            }

            path.append(i);
            used[i] = true;
            dfs(index + 1 ,path);
            return ;
        }
    }
}
// @lc code=end

