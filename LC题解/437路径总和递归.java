/*
437. 路径总和 III
给定一个二叉树，它的每个结点都存放着一个整数值。
找出路径和等于给定数值的路径总数。
路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
示例：
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1
返回 3。和等于 8 的路径有:
1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11
 */
/**
 * 递归法
 * 首先，最简单的子问题是什么呢？由于这道题是在树的框架下，我们最容易想到的就是遍历的终止条件：
 * if(root == null){
 *     return 0;
 * }
 * 接下来，我们来考虑再上升的一个层次，题目要求
 * 路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点） 。
 * 这就要求我们只需要去求三部分即可：
 *     以当前节点作为头结点的路径数量
 *     以当前节点的左孩子作为头结点的路径数量
 *     以当前节点的右孩子作为头结点啊路径数量
 * 将这三部分之和作为最后结果即可。
 * 最后的问题是：我们应该如何去求以当前节点作为头结点的路径的数量？
 * 这里依旧是按照树的遍历方式模板，每到一个节点让sum-root.val，
 * 并判断sum是否为0，如果为零的话，则找到满足条件的一条路径。
 */

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
    public int CountPath(TreeNode root, int sum){
        if(root == null){
            return 0;
        }
        sum = sum - root.val;
        int result  = sum == 0 ? 1 : 0;
        return result + CountPath(root.left, sum) + CountPath(root.right, sum);
    }
    public int pathSum(TreeNode root, int sum) {
        if(root == null){
            return 0;
        }
        int result = CountPath(root, sum);
        int a = pathSum(root.left, sum);
        int b = pathSum(root.right, sum);
        return result + a + b;
    }
}

