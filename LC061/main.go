package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	head := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	rotateRight(head, 1)
}
