# LC341 扁平化嵌套列表迭代器

## 方法一：深度优先搜索

### 思路

<u>嵌套的整型列表实际上是一个树形结构，树上的叶子节点对应一个整数，非叶子结点对应一个列表。</u>在这棵树上深度优先搜索的顺序就是迭代器遍历的顺序。

可以先遍历整个嵌套列表，将所有整数存入一个数组，然后遍历该数组从而实现next和hasNext方法。

```go
type NestedIterator struct {
	vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var vals []int
	var dfs func([]*NestedInteger)
	dfs = func(nesredList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() {
				vals = append(vals, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}
	dfs(nestedList)
	return & NestedIterator{vals}
}

func (this *NestedIterator) Next() int {
	val := this.vals[0]
	this.vals = this.vals[1:]
	return val
}

func (this *NestedIterator) HasNext() bool {
	return len(this.vals) > 0
}
```

## 方法二：栈

### 思路

我们可以用一个栈来代替方法一中的递归过程

具体，用一个栈来维护深度优先搜索时，从根节点到当前节点路径上的所有节点。由于非叶节点对应一个列表，我们在栈中存储的是指向列表当前遍历的元素的指针（下标）。每次向下搜索时，取出列表的当前指针指向的元素并将其入栈，同时将该指针指向后移动一位。如此反复直到找到一个整数。循环时若栈顶指针指向了列表末尾，则将其从栈顶弹出。

### 代码

```go
type NestedIterator struct {
    vals []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    var vals []int
    var dfs func([]*NestedInteger)
    dfs = func(nestedList []*NestedInteger) {
        for _, nest := range nestedList {
            if nest.IsInteger() {
                vals = append(vals, nest.GetInteger())
            } else {
                dfs(nest.GetList())
            }
        }
    }
    dfs(nestedList)
    return &NestedIterator{vals}
}

func (it *NestedIterator) Next() int {
    val := it.vals[0]
    it.vals = it.vals[1:]
    return val
}

func (it *NestedIterator) HasNext() bool {
    return len(it.vals) > 0
}
```

