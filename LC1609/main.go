package main

import "math"

func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	for level := 0; len(q) > 0; level++ {
		prev := 0
		if level%2 == 1 {
			prev = math.MaxInt32
		}
		size := len(q)
		for _, node := range q {
			val := node.Val
			if val%2 == level%2 || level%2 == 0 && val <= prev || level%2 == 1 && val >= prev {
				return false
			}
			prev = val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[size:]
	}
	return true
}
