# LC002 两数之和

## 方法：模拟

### 思路与算法

由于输入的两个链表都是逆序存储数字位数，所以同一位置的数字可以直接相加。

同时遍历两个链表，逐位计算他们的和，并与当前位置的进位值相加。具体而言：当前两个链表相应位置的数字为n1, n2，进位值设定为carry，则他们的和为n1+n2+carry，其中，单独设定一个答案链表，答案链表相应位置的数字为(n1+n2+carry)%10，而新的进位值为(n1+n2+carry)/10.

如果两个链表的长度不同，则可以认为短的链表后面有若干个0.

如果链表遍历结束，进位值carry>0，还需要在答案链表的后面附加一个额外节点，节点值为carry。

### 代码

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	// 同时遍历l1和l2链表
	for l1 != nil || l2 != nil {
		// 给n1，n2赋值
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		// 计算sum和carry
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			// 第一次头结点为空
			head = &ListNode{Val: sum}
			tail = head
		} else {
			// 由于头结点定义的是地址，所以尾节点后面衔接后，头结点对应也会增加节点
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		// 最后一次如果carry值不为0，那在后面增加一个额外节点
		if carry > 0 {
			tail.Next = &ListNode{Val: carry}
		}
	}
	return
}
```



### *Tips: Golang函数定义（支持命名返回值）*

```go
func add(x,y int) (sum int){    //命名返回值
    sum = x + y
    // 此处直接返回sum值
    return
}
```

