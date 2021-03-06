#### [173. 二叉搜索树迭代器](https://leetcode-cn.com/problems/binary-search-tree-iterator/)

实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 `next()` 将返回二叉搜索树中的下一个最小的数。

 

**示例：**

**![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/25/bst-tree.png)**

```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false
```

 

**提示：**

- `next()` 和 `hasNext()` 操作的时间复杂度是 O(1)，并使用 O(*h*) 内存，其中 *h* 是树的高度。
- 你可以假设 `next()` 调用总是有效的，也就是说，当调用 `next()` 时，BST 中至少存在一个下一个最小的数。

### 思路：

1. next()返回最小值，需要用slice存储中序遍历结果
2. 判断slice长度判断hasnext()

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
   Nums []int
   root *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
   tmpNums := make([]int, 0)
  
   var inorder func(node *TreeNode)
   inorder = func (node *TreeNode) {
      if node == nil {
         return
      }
      inorder(node.Left)
      tmpNums = append(tmpNums, node.Val)
      inorder(node.Right)
   }
   inorder(root)
   return BSTIterator{
      Nums: tmpNums,
      root: root,
   }
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
   ans := 0
   if this.HasNext() {
      ans = this.Nums[0]
      this.Nums = this.Nums[1 :]
   }
   return  ans
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
   if len(this.Nums) > 0 {
      return true
   }
   return false
}
```

