/*
 * @lc app=leetcode.cn id=201 lang=java
 *
 * [201] 数字范围按位与
 */

// @lc code=start
class Solution {
    public int rangeBitwiseAnd(int m, int n) {
        int shift = 0;

        while(m < n){
            m >>= 1;
            n >>= 1;
            ++shift;
        }

        return m << shift;
    }
}
// @lc code=end

