import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=5490 lang=java
 *
 * [5490] 吃掉 N 个橘子的最少天数
 */

// @lc code=start
class Solution {
    HashMap<Integer, Integer> map = new HashMap<>();

    public int minDays(int n) {
        if (n <= 1)return n;

        if (map.containsKey(n)){
            return map.get(n);
        }

        int min  = Integer.MAX_VALUE;
        min = Math.min(min, minDays(n/2) + n%2);
        min = Math.min (min, minDays(n/3) + n%3);

        min  +=1;
        map.put(n, min);

        return min;
    }
}
// @lc code=end
