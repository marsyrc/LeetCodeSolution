// 剑指 Offer 07. 重建二叉树

// 输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

 

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]

// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode {preorder[0], nil, nil}
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1 : len(inorder[:i]) + 1], inorder[:i])
    root.Right = buildTree(preorder[len(inorder[:i]) + 1 :], inorder[i + 1:])
    return root
}