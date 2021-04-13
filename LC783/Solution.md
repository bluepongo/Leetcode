# LC783 二叉搜索树节点最小距离

## 方法一：中序遍历

### 思路与算法

本体要求二叉搜索树任意两节点差的最小值，而我们知道二叉搜索树有个性质为二叉搜索树中序遍历得到的值的序列是递增有序的，因此我们只要得到中序遍历后的值序列就能用其他方法来解决

朴素的方法是经过一次中序遍历，将值保存在一个数组中再进行遍历求解，我们也可以在中序遍历的过程中用pre变量保存前驱结点的值，这样能边遍历边更新答案，不用再显式创建数组保存，但是要注意pre的初始值需要设置成任意负数标记开头

### 代码

```go
func minDiffInBST(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
```

