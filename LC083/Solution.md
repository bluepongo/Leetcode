# LC083 删除排序链表中的重复元素

## 方法一：一次遍历

### 思路与算法

由于给定的链表是排好序的，因此重复的元素在链表中的位置是连续的，所以只需要对链表进行一次遍历，就可以删除重复的元素

首先，我们从指针cur指向链表头结点，开始对链表进行遍历。如果当前cur和cur.next对应的元素相同，那么就将cur.next从链表中移除，否则就说明链表中不存在其他元素相同的节点，因此可以将cur指向cur.next

遍历完整个链表以后，我们返回链表的头节点

### 细节

遍历到链表的最后一个节点时，cur.next为空节点，如果不判断，访问cur.next对应的元素会产生运行错误。所以我们只需要遍历到链表的最后一个节点，而不需要遍历完整个链表

### 代码

```go
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// 删除重复值的节点
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
```

