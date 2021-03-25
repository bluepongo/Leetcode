# LC082 删除排序链表中的重复元素II

## 方法一：一次遍历

### 思路与算法

给定的链表已经排好序，所以重复的元素在链表中出现的位置是连续的，所以只需要进行一次遍历删除重复元素。由于头结点可能会被删除，所以构建一个虚拟节点指向链表头结点

具体地，从指针cur指向链表的虚拟节点，开始遍历。如果cur.next和cur.next.next对应的元素相同，就需要将cur.next以及后面拥有相同元素值的链表节点全部删除。记录下这个重复的具体值x，随后根据x的值不断将cur.next从链表中移除(让cur.next指向cur.next.next)，直到cur.next为空的或者元素值不为x为止。此时，将链表中所有元素为x的节点全部删除。

如果当前cur.next与cur.next.next对应的元素不相同，那么说明链表中只有一个元素值为cur.next的节点，那么可以将cur指向下一个cur.next

遍历完整个链表以后，返回链表虚拟节点的下一个节点dummy.next即可

### 细节

需要注意cur.next以及cur.next.next可能为空节点，需要额外加以判断

### 代码

```go
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 创建虚拟头结点
	dummy := &ListNode{0, head}
	// 改变cur也就是改变原数组，cur只是一个指针
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			// 保存重复值
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				// 删除重复节点
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
```

