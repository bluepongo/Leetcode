package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) (kth *ListNode) {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	for kth = head; n > k; n-- {
		kth = kth.Next
	}
	return
}
