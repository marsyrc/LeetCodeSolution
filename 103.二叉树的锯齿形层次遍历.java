import java.util.ArrayList;
import java.util.Stack;

import javax.swing.tree.TreeNode;

import jdk.nashorn.api.tree.Tree;

/*
 * @lc app=leetcode.cn id=103 lang=java
 *
 * [103] 二叉树的锯齿形层次遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode(int x) { val = x; } }
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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
    	List<List<Integer>> list = new ArrayList<>();
        if (root == null) {
        	return list;
        }
        //栈1来存储右节点到左节点的顺序
        Stack<TreeNode> stack1 = new Stack<>(); 
        //栈2来存储左节点到右节点的顺序
        Stack<TreeNode> stack2 = new Stack<>();
        
        //根节点入栈
    	stack1.push(root);
    	
    	//每次循环中，都是一个栈为空，一个栈不为空，结束的条件两个都为空
        while (!stack1.isEmpty() || !stack2.isEmpty()) {
        	List<Integer> subList = new ArrayList<>(); // 存储这一个层的数据
        	TreeNode cur = null;
        	
        	if (!stack1.isEmpty()) { //栈1不为空，则栈2此时为空，需要用栈2来存储从下一层从左到右的顺序
        		while (!stack1.isEmpty()) {	//遍历栈1中所有元素，即当前层的所有元素
        			cur = stack1.pop();
        			subList.add(cur.val);	//存储当前层所有元素
        			
        			if (cur.left != null) {	//左节点不为空加入下一层
        				stack2.push(cur.left);
        			}
        			if (cur.right != null) {	//右节点不为空加入下一层
        				stack2.push(cur.right);
        			}
        		}
        		list.add(subList);
        	}else {//栈2不为空，则栈1此时为空，需要用栈1来存储从下一层从右到左的顺序
        		while (!stack2.isEmpty()) {
        			cur = stack2.pop();
        			subList.add(cur.val);
        			
        			if (cur.right != null) {//右节点不为空加入下一层
        				stack1.push(cur.right);
        			}
        			if (cur.left != null) { //左节点不为空加入下一层
        				stack1.push(cur.left);
        			}
        		}
        		list.add(subList);
        	}
        }
        return list;
    }
}


// @lc code=end
