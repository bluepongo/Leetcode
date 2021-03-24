# LC456 132模式

## 方法一：枚举3

### 思路与算法

枚举3是容易想到并且也是最容易实现的。由于3是模式中的最大值，并且其出现在1和2的中间，因此我们只需要从左到右枚举3的下标j，那么：

- 由于1是模式中的最小值，因此在枚举j的同时，维护数组a中最左侧元素a[0..j-1]的最小值，即为1对应的的元素a[i]。需要注意，只有a[i]<a[j]时，a[i]才能作为1对应的元素
- 由于2是模式中的次小值，因此我们可以使用一个有序集合（平衡树）维护数组a中右侧元素a[j+1..n-1]中的所有值。当我们确定了a[i]和a[j]之后，只需要在平衡树中查询：严格比a[i]大的最小元素，即为a[k]。需要注意的是，只有a[k]<a[j]时，a[k]才能作为3对应的元素

### 代码

```go
type node struct {
	// 两个叶子节点
	ch       [2]*node
	priority int
	val      int
	cnt      int
}

// 比较节点值和b，b小返回0，b大返回1
func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

// 建立一个横二叉树
type treap struct {
	root *node
}

// 新增一个节点到树上，下划线为了区分私有方法，没有实际意义
func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._put(o.ch[d], val)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.cnt++
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.cnt > 1 {
		o.cnt--
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

// 寻找t中比val大的最小节点值
func (t *treap) upperBound(val int) (ub *node) {
	for o := t.root; o != nil; {
		if o.cmp(val) == 0 {
			ub = o
			o = o.ch[0]
		} else {
			o = o.ch[1]
		}
	}
	return
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// 数组维护左边数组最小值
	leftMin := nums[0]
	// 有序数组（平衡二叉树）维护右侧元素所有值
	rights := &treap{}
	for _, v := range nums[2:] {
		rights.put(v)
	}

	// 开始枚举3的下标j
	for j := 1; j < n-1; j++ {
		if nums[j] > leftMin {
			ub := rights.upperBound(leftMin)
			// 如果右边数组有比左侧最小值大的数，并且ub的值比nums[j]的小，返回true
			if ub != nil && ub.val < nums[j] {
				return true
			}
		} else {
			// 否则，让最小值来到j处
			leftMin = nums[j]
		}
		// 右边平衡二叉树删除第j+1个节点
		rights.delete(nums[j+1])
	}

	return false
}
```

## 方法二：枚举1

### 思路与算法

如果我们从左到右枚举1的下标i，那么j，k的下表范围都是减少的，不利于对他们进行维护。所有考虑从右到左枚举i

132模式中，如果1<2并且2<3，呢么根据传递性，1<3，所以可以使用以下方法进行维护：

- 使用一种数据结构维护所有遍历过的元素，作为2的候选元素。每当我们遍历到一个新的元素，就将其加入数据结构中。
- 遍历到一个新的元素的同时，可以考虑其是否可以作为3。如果作为3，那么<u>数据结构中所有严格小于他的元素都可以作为2</u>，我们将这些元素全部从数据结构中移除，并且<u>使用一个变量维护所有被移除的元素的最大值</u>。这些被移除的元素都是可以真正作为2的，并且元素的值越大，之后找到1的机会也就越大。

这个数据结构需要满足：

- 支持添加一个元素
- 支持移除所有严格小于给定阈值的所有元素
- 上面两步是依次进行的，也就是说我们先用给定的阈值移除元素，再将阈值加入数据结构中

使用「单调栈」。单调栈中，栈底到栈顶的元素是严格单调递减的。给定阈值阈值x时，我们只需要不断弹出栈顶的元素，直到栈为空或者x严格小于栈顶元素。此时再将x入栈，九尾狐了栈的单调性。

所以我们使用单调栈作为维护2的数据结构，给出下列算法：

- 用单调栈维护所有可以作为2的候选元素。初始时，单调栈中只有唯一的元素a[n-1]。我们还需要用一个变量max_k记录所有可以真正作为2的元素的最大值
- 然后我们从n-2开始从右向左枚举元素a[i]：
  - 首先判断a[i]是否可以作为1.如果a[i]<max_k,就可以作为1，就找到了一组满足132模式的三元组
  - 判断a[i]是否可以作为3，来找出哪些可以真正作为2的元素。将a[i]不断地与单调栈顶的元素进行比较，如果a[i]比较大，那么栈顶元素可以真正作为2，将其弹出并且更新max_k
  - 最后将a[i]作为2的候选元素放入单调栈中。这里可以进行一个优化，如果a[i]<= max_k，那么我们没有必要将a[i]放入栈中，因为即使他在未来被弹出，也不会将max_k更新为更大的值
- 枚举完所有元素之后，如果仍然没有找到满足132的三元组，就不存在

### 代码

```go
func find132pattern002(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

  	
	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}

	return false
}
```

