/*
 * @lc app=leetcode.cn id=40 lang=java
 *
 * [40] 组合总和 II
 */

// @lc code=start
class Solution {

    private int n;

    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        List<List<Integer>> ans = new ArrayList<List<Integer>>();
        List<Integer> path = new ArrayList<Integer>();
        Arrays.sort(candidates);
        n = candidates.length;
        backTrace(candidates,target,ans,path,0);
        return ans;
    }

    private void backTrace(int[] candidates, int target, List<List<Integer>> ans, List<Integer> path, int index) {
        //终止条件
        if (target == 0) {
            ans.add(new ArrayList<>(path));
            return;
        }

        for (int i = index; i < n; i++) {
            //合法性检查
            if (target < candidates[i]) {
                continue;
            }
            if (i > index && candidates[i - 1] == candidates[i]) {
                continue;
            }

            //选择、递归、撤销
            path.add(candidates[i]);
            backTrace(candidates, target - candidates[i], ans, path, i + 1);
            path.remove(path.size() - 1);
        }
    }
}
// @lc code=end

