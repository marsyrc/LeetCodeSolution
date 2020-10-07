import java.util.ArrayList;
import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        int n = nums.length;
        Arrays.sort(nums);

        List<List<Integer>> ans = new ArrayList<List<Integer>>();

        // 枚举a
        for (int first = 0; first < n; first++) {
            //和上一次枚举的数字不同，避免重复
            if (first > 0 && nums[first] == nums[first - 1]) {
                continue;
            }

            int third  = n - 1;

            int  target = - nums[first];

            //枚举b
            for (int second = first + 1; second < n; second++) {
                //避免重复
                if (second > first + 1 && nums[second ] == nums[second -1 ]){
                    continue;
                }

                //保证b在c左边
                while (second < third && nums [second ]  + nums[third] > target ){
                    third--;
                }

                if (second == third ){
                    break ;
                }

                if (nums[second] + nums[third] == target){
                    List<Integer> list = new ArrayList<Integer>();
                    list.add(nums[first]);
                    list.add(nums[second]);
                    list.add(nums[third]);
                    ans.add(list);
                }
            }
        }

        return ans;
    }
}
// @lc code=end
