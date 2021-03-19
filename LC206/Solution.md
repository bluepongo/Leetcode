# LC206 反转链表

## 方法：迭代

假设链表为1→2→3→∅，我们想要把它改成 ∅←1←2←3。

在遍历链表时，将当前节点 `next`指针改为指向前一个节点。

由于节点没有引用其前一个节点，因此必须事先存储其前一个节点。

在更改引用之前，还需要存储后一个节点。最后返回新的头引用

## 代码

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
      	// 分离curr.Next
        next := curr.Next
        // 反转箭头指向
        curr.Next = prev
        // 重新定义prev和curr
        prev = curr
        curr = next
    }
    return prev
}
```

