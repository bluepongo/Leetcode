# LC094 二叉树的中序遍历

## 方法一：递归

### 思路与算法

二叉树的中序遍历：访问左子树->根节点->右子树的方式遍历这棵树，而在访问左子树或者右子树的时候，我们按照同样的方式遍历，直到遍历完整棵树。因此整个遍历过程天然具有递归的性质，可以直接用递归函数模拟这个过程

定义inorder(root)表示当前遍历到root节点的答案，按照定义，我们只要递归调用inorder(root.left)来遍历root节点的左子树，然后将root节点的值加入答案，再递归调用inorder(root.right)来遍历root节点的右子树即可，递归的终止条件是碰到空节点

### 代码

```go
func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
```

## 方法二：迭代

### 思路与算法

方法一的函数也可以用迭代的方式实现，两种方式是层架的，区别是，递归的时候隐式维护了一个栈，迭代则是显示地把这个栈模拟出来

### 代码

```go
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
```

## *方法三：Morris中序遍历

### 思路与算法

Morris遍历算法是另一种遍历二叉树的方法，能将非递归的中序遍历空间复杂度降为O(1)

Morris遍历算法步骤如下（假设当前遍历到的节点为x）：

- 如果x无左孩子，先将x加入答案数组，再访问x的右孩子（x=x.right）
- 如果x有左孩子，找到x左子树上最右的节点（<u>左子树中序遍历的最后一个节点，x在中序遍历中的前驱节点</u>）记为predecessor。再根据predecessor的右孩子是否为空，进行如下操作
  - 如果predecessor的右孩子为空，则将其右孩子指向x，然后访问x的左孩子，即x=x.left
  - 如果predecessor的右孩子不为空，则此时其右孩子指向x，说明已经遍历完了x的左子树，将predecessor的右孩子置空，将x的值加入答案，再访问x的右孩子（x.right）
- 重复上述操作，直至访问完整棵树

### 代码

```go
func inorderTraversal(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}
```

