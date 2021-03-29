# LC173 二叉搜索迭代器

## 方法一：扁平化

### 思路

可以直接对二叉搜索树做一次完全的递归遍历，获取到中序遍历的全部结果保存自数组中。随后，我们利用得到的数组本身来实现迭代器

### 代码

```go
type BSTIterator struct {
    arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
    it.inorder(root)
    return
}

func (it *BSTIterator) inorder(node *TreeNode) {
    if node == nil {
        return
    }
    // 递归将中序遍历的值存入迭代器
    it.inorder(node.Left)
    it.arr = append(it.arr, node.Val)
    it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
    val := it.arr[0]
    it.arr = it.arr[1:]
    return val
}

func (it *BSTIterator) HasNext() bool {
    return len(it.arr) > 0
}
```

## 方法二：迭代

### 思路

利用栈这一个数据结构，通过迭代的方式对二叉树进行中序遍历。此时，不用预先算出中序遍历的全部结果，只需要实时维护当前栈的情况即可

### 代码

```go
type BSTIterator struct {
    stack []*TreeNode
    cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{cur: root}
}

func (it *BSTIterator) Next() int {
    for node := it.cur; node != nil; node = node.Left {
        it.stack = append(it.stack, node)
    }
    it.cur, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
    val := it.cur.Val
    it.cur = it.cur.Right
    return val
}

func (it *BSTIterator) HasNext() bool {
    return it.cur != nil || len(it.stack) > 0
}
```

