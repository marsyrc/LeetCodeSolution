/*
 * @lc app=leetcode.cn id=257 lang=java
 *
 * [257] 二叉树的所有路径
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
    public List<String> binaryTreePaths(TreeNode root) {
        List<String> paths = new ArrayList<String>();

        if (root == null) {
            return paths;
        }

        Queue<TreeNode> nodeQueue = new LinkedList<TreeNode>();
        Queue<String> pathQuene = new LinkedList<String>();

        nodeQueue.offer(root);
        pathQuene.offer(Integer.toString(root.val));

        while (!nodeQueue.isEmpty()) {
            TreeNode node = nodeQueue.poll();
            String path = pathQuene.poll();
            if (node.left == null && node.right == null)paths.add(path);
            if (node.left != null){
                nodeQueue.offer(node.left);
                pathQuene.offer(new StringBuffer(path).append("->").append(node.left.val).toString());
            }

            if (node.right != null){
                nodeQueue.offer(node.right);
                pathQuene.offer(new StringBuffer(path).append("->").append(node.right.val).toString());
            }
        }
        return paths;
    }
}
// @lc code=end

