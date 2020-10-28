import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=46 lang=java
 *
 * [46] 全排列
 */

// @lc code=start
class Solution {
    List<List<Integer>> res = new LinkedList<>();

    public List<List<Integer>> permute(int[] nums) {
        LinkedList<Integer> track = new LinkedList<>();
        backtrack(track, nums);
        return res;
    }

    void backtrack(LinkedList<Integer> track, int[] nums) {
        if (track.size() ==  nums.length){
            res.add(new LinkedList<>(track));
            return;
        }

        for (int i = 0; i <nums.length; ++i){
            if (track.contains(nums[i])){
                continue;
            }

            track.add(nums[i]);

            backtrack(track, nums);

            track.removeLast();
        }
    }
}

// @lc code=end
