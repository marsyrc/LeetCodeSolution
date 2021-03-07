import java.util.List;
import java.util.Stack;

/*
 * @lc app=leetcode.cn id=107 lang=java
 *
 * [107] 二叉树的层次遍历 II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
/**
 * 将正常的层序遍历结果放入 stack 中，再出栈放入结果ans
 */
class Solution {
    public List<List<Integer>> levelOrderBottom(TreeNode root) {

        List<List<Integer>> ans = new ArrayList<>();

        Queue<TreeNode> queue = new ArrayDeque<>();

        Stack<List<Integer>> stack = new Stack<>();

        if (root != null)
            queue.add(root);

        while (!queue.isEmpty()) {
            List<Integer> Level = new ArrayList<>();
            int n = queue.size();

            for (int i = 0; i < n; i++) {
                TreeNode currNode = queue.poll();
                Level.add(currNode.val);
                if (currNode.left != null) {
                    queue.add(currNode.left);
                }
                if (currNode.right != null) {
                    queue.add(currNode.right);
                }
            }
            stack.add(Level);
        }

        while (!stack.isEmpty()) {
            ans.add(stack.pop());
        }
        return ans;
    }
}
// @lc code=end
