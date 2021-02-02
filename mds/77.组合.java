/*
 * @lc app=leetcode.cn id=77 lang=java
 *
 * [77] 组合
 */

// @lc code=start
class Solution {
    public List<List<Integer>> combine(int n, int k) {
        if(n < k || k <= 0)return null;
        
        List<List<Integer>> ans = new ArrayList<>();
        List<Integer> path = new ArrayList<>();
        dfs(n , k , 1 ,path, ans);
        return ans;
    }
    
    private void dfs (int n , int k , int index , List<Integer>path, List<List<Integer>> ans){
        if (path.size() == k){
            ans.add(new ArrayList<>(path));
            return;
        }

         /**
            i 从 index 开始，保证 在当前循环中，不会重复遍历 之前遍历过的数字
            n - (k - path.size()) + 1 表示 当前循环，起点所能达到的最大的数(保证最后一个数字加入时，正好能组成长度为k的结果)
         */
        for (int i = index; i <= n - (k - path.size())+1; i++) {


            path.add(i);

            dfs(n, k , i+1 , path,ans);

            path.remove(path.size()-1);

        }
    }
}
// @lc code=end

