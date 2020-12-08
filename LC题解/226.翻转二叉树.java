import jdk.nashorn.api.tree.Tree;

/*
 * @lc app=leetcode.cn id=226 lang=java
 *
 * [226] 翻转二叉树
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
class Solution {
    public TreeNode invertTree(TreeNode root) {
        if(root == null)return null;

        TreeNode newRight = invertTree(root.left); 
        TreeNode newLeft = invertTree(root.right);

        root.left = newLeft;
        root.right = newRight;

        return root;
    }
}
// @lc code=end

