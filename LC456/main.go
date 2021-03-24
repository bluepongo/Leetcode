package main

import (
	"fmt"
	"math"
	"math/rand"
)

type node struct {
	// 两个叶子节点
	ch       [2]*node
	priority int
	val      int
	cnt      int
}

// 比较节点值和b，b小返回0，b大返回1
func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

// 建立一个横二叉树
type treap struct {
	root *node
}

// 新增一个节点到树上，下划线为了区分私有方法，没有实际意义
func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val, cnt: 1}
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._put(o.ch[d], val)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	} else {
		o.cnt++
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if o == nil {
		return nil
	}
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.cnt > 1 {
		o.cnt--
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

// 寻找t中比val大的最小节点值
func (t *treap) upperBound(val int) (ub *node) {
	for o := t.root; o != nil; {
		if o.cmp(val) == 0 {
			ub = o
			o = o.ch[0]
		} else {
			o = o.ch[1]
		}
	}
	return
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	// 数组维护左边数组最小值
	leftMin := nums[0]
	// 有序数组（平衡二叉树）维护右侧元素所有值
	rights := &treap{}
	for _, v := range nums[2:] {
		rights.put(v)
	}

	// 开始枚举3的下标j
	for j := 1; j < n-1; j++ {
		if nums[j] > leftMin {
			ub := rights.upperBound(leftMin)
			// 如果右边数组有比左侧最小值大的数，并且ub的值比nums[j]的小，返回true
			if ub != nil && ub.val < nums[j] {
				return true
			}
		} else {
			// 否则，让最小值来到j处
			leftMin = nums[j]
		}
		// 右边平衡二叉树删除第j+1个节点
		rights.delete(nums[j+1])
	}

	return false
}

func find132pattern002(nums []int) bool {
	n := len(nums)
	candidateK := []int{nums[n-1]}
	maxK := math.MinInt64

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}
	return false
}

func main() {
	n := []int{-1, 4, 3, 0}
	res := find132pattern002(n)
	fmt.Println(res)

}
