package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	deleteDuplicates(head)
}
