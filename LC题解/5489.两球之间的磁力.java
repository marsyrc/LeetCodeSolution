/*
 * @lc app=leetcode.cn id=5489 lang=java
 *
 * [5489] 两球之间的磁力
 */

// @lc code=start

/**
 * 
 * @param position
 * @param distance
 * @param m
 * @return
 */

 /*
 二分查找加贪心搜索
  */
class Solution {
    public int maxDistance(int[] position, int m) {
        Arrays.sort(position);
        int high = (position[position.length - 1] - position[0]) / (m - 1);
        int low = 1;
        int ans = 1;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (check(position,mid, m)){
                ans = mid;
                low = mid + 1;
            }else {high = mid -1;}
        }
        return ans;
    }

    public boolean check(int[] position, int distance, int m) {
        int count = 1;
        int i = 0;
        for (int j = 1; j < position.length; j++) {
            if ((position[j] - position[i]) >= distance) {
                i = j;
                count++;
                if (count >= m) return true;
            }
        }
        return false;
    }
}
// @lc code=end

