力扣501

**我们用 Morris 中序遍历的方法把中序遍历的空间复杂度优化到 \*O(1)\*。** 我们在中序遍历的时候，一定先遍历左子树，然后遍历当前节点，最后遍历右子树。在常规方法中，我们用递归回溯或者是栈来保证遍历完左子树可以再回到当前节点，但这需要我们付出额外的空间代价。我们需要用一种巧妙地方法可以在 *O(1)* 的空间下，遍历完左子树可以再回到当前节点。我们希望当前的节点在遍历完当前点的前驱之后被遍历，我们可以考虑修改它的前驱节点的 right 指针。当前节点的前驱节点的 right  指针可能本来就指向当前节点（前驱是当前节点的父节点），也可能是当前节点左子树最右下的节点。如果是后者，我们希望遍历完这个前驱节点之后再回到当前节点，可以将它的 right 指针指向当前节点。

Morris 中序遍历的一个重要步骤就是寻找当前节点的前驱节点，并且 Morris 中序遍历寻找下一个点始终是通过转移到 right  指针指向的位置来完成的。

- 如果当前节点没有左子树，则遍历这个点，然后跳转到当前节点的右子树。
- 如果当前节点有左子树，那么它的前驱节点一定在左子树上，我们可以在左子树上一直向右行走，找到当前点的前驱节点。
  - 如果前驱节点没有右子树，就将前驱节点的 right 指针指向当前节点。这一步是为了在遍历完前驱节点后能找到前驱节点的后继，也就是当前节点。
  - 如果前驱节点的右子树为当前节点，说明前驱节点已经被遍历过并被修改了 right  指针，这个时候我们重新将前驱的右孩子设置为空，遍历当前的点，然后跳转到当前节点的右子树。

因此我们可以得到这样的代码框架：

```cpp
TreeNode *cur = root, *pre = nullptr;
while (cur) {
    if (!cur->left) {
        // ...遍历 cur
        cur = cur->right;
        continue;
    }
    pre = cur->left;
    while (pre->right && pre->right != cur) {
        pre = pre->right;
    }
    if (!pre->right) {
        pre->right = cur;
        cur = cur->left;
    } else {
        pre->right = nullptr;
        // ...遍历 cur
        cur = cur->right;
    }
}
```

```Golang
func findMode(root *TreeNode) (answer []int) {
    var base, count, maxCount int
    update := func(x int) {
        if x == base {
            count++
        } else {
            base, count = x, 1
        }
        if count == maxCount {
            answer = append(answer, base)
        } else if count > maxCount {
            maxCount = count
            answer = []int{base}
        }
    }
    cur := root
    for cur != nil {
        if cur.Left == nil {
            update(cur.Val)
            cur = cur.Right
            continue
        }
        pre := cur.Left
        for pre.Right != nil && pre.Right != cur {
            pre = pre.Right
        }
        if pre.Right == nil {
            pre.Right = cur
            cur = cur.Left
        } else {
            pre.Right = nil
            update(cur.Val)
            cur = cur.Right
        }
    }
    return
}
```

**复杂度分析**

- 时间复杂度：*O(n)*。每个点被访问的次数不会超过两次，故这里的时间复杂度是 *O(n)*。
- 空间复杂度：*O(1)*。使用临时空间的大小和输入规模无关。