import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=214 lang=java
 *
 * [214] 最短回文串
 */

// @lc code=start
// 我们使用 KMP 字符串匹配算法来找出这个最长的前缀回文串。
// 具体地，记 s^ 为 s的反序，由于 s1 是 s的前缀，
// 那么 s^1​ 就是 s^ 的后缀。
// 考虑到 s1 是一个回文串，
//因此 s1=s^1 ​，s1​ 同样是 s^的后缀。
// 这样一来，我们将 s 作为模式串，s^作为查询串进行匹配。
//当遍历到 s^的末尾时，如果匹配到 s中的第 i 个字符，
//那么说明 s 的前 i 个字符与 s^ 的后 i个字符相匹配（即相同），
//s 的前 i 个字符对应 s1,s^ 的后 i 个字符对应 s^1​，由于 s1=s^1​，因此 s1​ 就是一个回文串。
class Solution {
    public String shortestPalindrome(String s) {
        String ss = s + '#' + new StringBuilder(s).reverse();
        int max = getLastNext(ss);
        return new StringBuilder(s.substring(max)).reverse() + s;
    }
    
    //返回 next 数组的最后一个值
    public int getLastNext(String s) {
        int n = s.length();
        char[] c = s.toCharArray();
        int[] next = new int[n + 1];
        next[0] = -1;
        next[1] = 0;
        int k = 0;
        int i = 2;
        while (i <= n) {
            if (k == -1 || c[i - 1] == c[k]) {
                next[i] = k + 1;
                k++;
                i++;
            } else {
                k = next[k];
            }
        }
        return next[n];
    }

}
// @lc code=end

