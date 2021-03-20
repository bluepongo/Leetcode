# LC150 逆波兰表达式求值

## 方法一：栈

### 思路

逆波兰表达式严格遵循「从左到右」的运算。计算逆波兰表达式的值时，使用一个栈存储操作数，从左到右遍历逆波兰表达式，进行如下操作：

- 如果遇到操作数，则将操作数入栈；

- 如果遇到运算符，则将两个操作数出栈，其中先出栈的是右操作数，后出栈的是左操作数，使用运算符对两个操作数进行运算，将运算得到的新操作数入栈。


整个逆波兰表达式遍历完毕之后，栈内只有一个元素，该元素即为逆波兰表达式的值。

### 代码

```go
func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			// 如果是数字，入栈
			stack = append(stack, val)
		} else {
			// 如果是符号，两个数字依次出栈
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
```

## 方法二：数组模拟栈

方法一使用栈存储操作数。也可以使用一个数组模拟栈操作。

### 思路

如果使用数组代替栈，则需要预先定义数组的长度。对于长度为n的逆波兰表达式，显然栈内元素个数不会超过n，但是将数组的长度定义为n仍然超过了栈内元素个数的上界。

对于一个有效的逆波兰表达式，其长度n一定是奇数，且操作数的个数一定比运算符的个数多1个，即包含$\frac{n+1}{2}$个操作数和$\frac{n-1}{2}$个运算符。考虑遇到操作数和运算符时，栈内元素个数分别会如何变化：

- <u>如果遇到操作数，则将操作数入栈，因此栈内元素增加1个；</u>
- <u>如果遇到运算符，则将两个操作数出栈，然后将一个新操作数入栈，因此栈内元素先减少2个再增加1个，结果是栈内元素减少1个。</u>

由此可以得到操作数和运算符与栈内元素个数变化的关系：<u>遇到操作数时，栈内元素增加 1个；遇到运算符时，栈内元素减少1个。</u>

最坏情况下，$\frac{n+1}{2}$个操作数都在表达式的前面，$\frac{n-1}{2}$个运算符都在表达式的后面，此时栈内元素最多为$\frac{n+1}{2}$个。在其余情况下，栈内元素总是少于$\frac{n+1}{2}$个。因此，在任何情况下，栈内元素最多可能有$\frac{n+1}{2}$个，将数组的长度定义为$\frac{n+1}{2}$即可。

### 实现方法

具体实现方面，创建数组 `stack`模拟栈，数组下标0的位置对应栈底，定义 `index`表示栈顶元素的下标位置，初始时栈为空，`index-1`。当遇到操作数和运算符时，进行如下操作：

- 如果遇到操作数，则将`index`的值加1，然后将操作数赋给`stack[index]`；

- 如果遇到运算符，则将`index`的值减1，此时`stack[index]`和`stack[index+1]`的元素分别是左操作数和右操作数，使用运算符对两个操作数进行运算，将运算得到的新操作数赋给`stack[index]`。


整个逆波兰表达式遍历完毕之后，栈内只有一个元素，因此 `index=0`，此时 `stack[index]`即为逆波兰表达式的值。

### 代码

```go
func evalRPN002(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			// 如果是操作数，index++，把val赋值给stack[index]
			index++
			stack[index] = val
		} else {
			// 如果是操作符，index--，令stack[index]和stack[index+1]进行操作
			index--
			switch token {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			case "/":
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[index]
}
```

