/*
 * @lc app=leetcode.cn id=39 lang=java
 *
 * [39] 组合总和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> ans = new ArrayList<List<Integer>>();
        List<Integer> path = new ArrayList<Integer>();
        dfs(candidates,target,ans,path,0);
        return ans;

    }

    //index为目前处理的索引
    private void dfs(int[] candidates, int target, List<List<Integer>> ans, List<Integer> path, int index) {
        if (index == candidates.length) {
            return;
        }

        if(target == 0){
            ans.add(new ArrayList<Integer>(path));
            return;
        }

        //选择直接跳过这个点
        dfs(candidates,target,ans,path,index+ 1);
        //如果合法，选择加入这个点
        if ((target - candidates[index]) >= 0){
            path.add(candidates[index]);
            dfs(candidates,target- candidates[index],ans,path,index);
            path.remove(path.size() - 1);
        }
    }
}
// @lc code=end

