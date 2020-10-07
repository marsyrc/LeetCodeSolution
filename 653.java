/*
653. 两数之和 IV - 输入 BST
给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
案例 1:
输入:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 9
输出: True
案例 2:
输入:
    5
   / \
  3   6
 / \   \
2   4   7
Target = 28
输出: False
 */

import java.util.HashSet;
import java.util.Set;

/**
 * 哈希表
 */
class Solution {
    public boolean findTarget(TreeNode root, int k) {
        Set <Integer> set = new HashSet();
        return find(root, k , set);
    }

    public boolean find(TreeNode root, int k, Set<Integer> set) {
        if (root == null) {
            return false;
        }
        if (set.contains(k - root.val)) {
            return true;
        }
        set.add(root.val);
        return find(root.left, k , set) || find(root.right, k , set);
    }
}