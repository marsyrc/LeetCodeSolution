/*
95. 不同的二叉搜索树 II
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
 */

import java.util.LinkedList;
import java.util.List;

/**
 * 递归
 */
class Solution {
    public List<TreeNode> generateTrees(int n) {
        if (n == 0) {
            return new LinkedList<TreeNode>();
        }
        return generateTrees(1, n);
    }

    public List<TreeNode> generateTrees(int start, int end) {
        List<TreeNode> allTrees = new LinkedList<TreeNode>();

        if (start > end) {
            allTrees.add(null);
            return allTrees;
        }

        for (int i = start; i <= end; ++i) {
            List<TreeNode> leftTrees = generateTrees(start, i - 1);
            List<TreeNode> rightTrees = generateTrees(i + 1, end);


            // 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
            for (TreeNode left : leftTrees) {
                for (TreeNode right : rightTrees) {
                    TreeNode current = new TreeNode(i);
                    current.left = left;
                    current.right = right;
                    allTrees.add(current);
                }
            }
        }
        return allTrees;
    }
}