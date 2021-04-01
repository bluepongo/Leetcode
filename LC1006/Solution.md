# LC1006 笨阶乘

## 方法一：使用栈模拟

### 思路

「笨阶乘」没有显式括号，运算优先级是先「乘除」后「加减」。我们可以从N开始，枚举N-1，N-2一直到1，枚举这些数的时候，认为他们之前的操作符按照「乘」「除」「加」「减」交替进行

- 出现乘法、除法的时候，可以把栈顶元素取出，与当前的N进行乘法、除法运算，并将运算结果重新压入栈中
- 出现加法、减法的时候，把减法视为加上一个数的相反数然后压入栈，等待以后遇见「乘」「除」法的时候取出

最后栈中的元素累加即为答案

### 代码

```go
func clumsy(N int) (ans int) {
	stk := []int{N}
	N--

	index := 0 // 用于控制乘、除、加、减
	for N > 0 {
		switch index % 4 {
		case 0:
			stk[len(stk)-1] *= N
		case 1:
			stk[len(stk)-1] /= N
		case 2:
			stk = append(stk, N)
		default:
			stk = append(stk, -N)
		}
		N--
		index++
	}

	// 累加栈中数字
	for _, v := range stk {
		ans += v
	}
	return
}
```

