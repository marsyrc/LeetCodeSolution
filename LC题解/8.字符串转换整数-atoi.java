/*
 * @lc app=leetcode.cn id=8 lang=java
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
class Solution {
    public int myAtoi(String str) {
        char[] chars = str.toCharArray();
        int n = chars.length;
        int i = 0; // 索引
        while (i < n && chars[i] == ' ') { // 去掉前导空格
            i++;
        }

        if (i == n) {
            return 0;
        }

        boolean negative = false;
        if (chars[i] == '-') {
            negative = true;
            i++;
        } else if (chars[i] == '+') {
            i++;
        } else if (!Character.isDigit(chars[i])) {
            return 0;
        }

        int ans = 0;
        while (i < n && Character.isDigit(chars[i])) {
            int digit = chars[i] - '0';
            if (ans > (Integer.MAX_VALUE - digit) / 10) {  //判断是否越界
                return negative ? Integer.MIN_VALUE : Integer.MAX_VALUE;
            }
            ans = ans * 10 + digit;
            i++;
        }
        return negative ? -ans : ans;
    }
}
// @lc code=end
