/*
 * @lc app=leetcode.cn id=29 lang=java
 *
 * [29] 两数相除
 */

// @lc code=start
class Solution {
    public int divide(int dividend, int divisor) {
        if (dividend == 0) return 0;

        if (divisor == 1) return dividend;

        if (divisor == -1) {
            if (dividend > Integer.MIN_VALUE) {
                return -dividend;
            } else return -(Integer.MIN_VALUE + 1);  // -min_value / -1 需要返回 max_value
        }

        long a = dividend;
        long b = divisor;
        int sign = 1;

        if (a * b < 0) {
            sign = -1;
        }

        a = Math.abs(a);
        b = Math.abs(b);
        int ans = div(a,b);
        if (sign > 0 ){
            return ans;
        }
        return -ans;
    }

    private int div(long a, long b) {
        if (a < b ) return  0;
        long count = 1;
        long tb = b ;
        while(2 * tb <= a ){
            count += count;
            tb +=tb;
        }
        return (int) count + div(a - tb, b);
    }
}
// @lc code=end

