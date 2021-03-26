# LC061 旋转链表

## 方法一：闭合为环

### 思路及算法

记录给定链表的长度为n，向右移动次数k>=n的时候，只需要向右移动k mod n次即可。因为每n次移动都会让链表变为原状。这样我们就知道，<u>新链表的最后一个节点为原链表的第（n-1)-(k mod n）个节点</u>（从0开始计数）

这样，我们先把给定的链表连接成环，然后从指定位置断开

具体代码中，我们首先计算出链表长度n，并找到链表的末尾节点，将其与头结点相连。得到了一个闭合为环的链表。然后我们找到新链表的最后一个节点（原链表的第（n-1）-（k mod n））个节点，将当前闭合为环的链表断开，得到结果

当链表长度不大于1，或者k为n的倍数时，新链表和原链表相同，不作处理

### 代码

```go
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	// 计算链表长度
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	// 从add位置处断开
	add := n - k%n
	if add == n {
		return head
	}
	// 首尾进行相连
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
```

