# LC092 反转链表II

## 方法一：穿针引线

我们以下图中黄色区域的链表反转为例。

![image.png](https://pic.leetcode-cn.com/1615105129-iUPoGi-image.png)

使用「206. 反转链表」的解法，反转 left 到 right 部分以后，再拼接起来。我们还需要记录 left 的前一个节点，和 right 的后一个节点。如图所示：

![image.png](https://pic.leetcode-cn.com/1615105150-pfWiGq-image.png)

### 算法步骤

第 1 步：先将待反转的区域反转；
第 2 步：把 pre 的 next 指针指向反转以后的链表头节点，把反转以后的链表的尾节点的 next 指针指向 succ。

![image.png](https://pic.leetcode-cn.com/1615105168-ZQRZew-image.png)

### 代码

```go
func reverseLinkedList(head *ListNode) {
    var pre *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
    // 因为头节点有可能发生变化，使用虚拟头节点可以避免复杂的分类讨论
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head

    pre := dummyNode
    // 第 1 步：从虚拟头节点走 left - 1 步，来到 left 节点的前一个节点
    // 建议写在 for 循环里，语义清晰
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }

    // 第 2 步：从 pre 再走 right - left + 1 步，来到 right 节点
    rightNode := pre
    for i := 0; i < right-left+1; i++ {
        rightNode = rightNode.Next
    }

    // 第 3 步：切断出一个子链表（截取链表）
    leftNode := pre.Next
    curr := rightNode.Next

    // 注意：切断链接
    pre.Next = nil
    rightNode.Next = nil

    // 第 4 步：同第 206 题，反转链表的子区间
    reverseLinkedList(leftNode)

    // 第 5 步：接回到原来的链表中
    pre.Next = rightNode
    leftNode.Next = curr
    return dummyNode.Next
}
```

## 方法二：一次遍历【穿针引线】反转链表（头插法）

方法一的缺点是：如果 left 和 right 的区域很大，恰好是链表的头节点和尾节点时，找到 left 和 right 需要遍历一次，反转它们之间的链表还需要遍历一次，虽然总的时间复杂度为 O(N)，但遍历了链表 2 次。

 ![image.png](https://pic.leetcode-cn.com/1615105232-cvTINs-image.png)

整体思想是：在需要反转的区间里，每遍历到一个节点，让这个新节点来到反转部分的起始位置。下面的图展示了整个流程。

![image.png](https://pic.leetcode-cn.com/1615105242-ZHlvOn-image.png)

下面我们具体解释如何实现。使用三个指针变量 pre、curr、next 来记录反转的过程中需要的变量，它们的意义如下：

- curr：指向待反转区域的第一个节点 left；
- next：永远指向 curr 的下一个节点，循环过程中，curr 变化以后 next 会变化；
- pre：永远指向待反转区域的第一个节点 left 的前一个节点，在循环过程中不变。

第 1 步，我们使用 ①、②、③ 标注「穿针引线」的步骤。

![image.png](https://pic.leetcode-cn.com/1615105296-bmiPxl-image.png)

### 操作步骤

- 先将 curr 的下一个节点记录为 next；
- 执行操作 ①：把 curr 的下一个节点指向 next 的下一个节点；
- 执行操作 ②：把 next 的下一个节点指向 pre 的下一个节点；
- 执行操作 ③：把 pre 的下一个节点指向 next

第 1 步完成以后「拉直」的效果如下：

![image.png](https://pic.leetcode-cn.com/1615105340-UBnTBZ-image.png)

第 2 步，同理。同样需要注意 **「穿针引线」操作的先后顺序**。

![image.png](https://pic.leetcode-cn.com/1615105353-PsCmzb-image.png)

第 2 步完成以后「拉直」的效果如下：

![image.png](https://pic.leetcode-cn.com/1615105364-aDIFqy-image.png)

第3步，同理。

![image.png](https://pic.leetcode-cn.com/1615105376-jIyGwv-image.png)

第 3 步完成以后「拉直」的效果如下：

![image.png](https://pic.leetcode-cn.com/1615105395-EJQnMe-image.png)

### 代码

```go
func reverseBetween(head *ListNode, left, right int) *ListNode {
    // 设置 dummyNode 是这一类问题的一般做法
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head
  	// 寻找pre，从虚拟头结点走left-1步，走到left节点的前一节点
    pre := dummyNode
    for i := 0; i < left-1; i++ {
        pre = pre.Next
    }
  	// 寻找cur，cur处于left节点处
    cur := pre.Next
    for i := 0; i < right-left; i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
    }
    return dummyNode.Next
}
```

