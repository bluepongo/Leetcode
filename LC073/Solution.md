# LC073 矩阵置零

## 方法一：使用标记数组

### 思路和算法

使用两个标记（row，col）分别记录每一行和每一列是否有零出现

具体，首先遍历该数组一次，如果某个元素为0，<u>那么就将这个元素锁在的行和列对应的标记位置设置为true</u>，最后再次遍历该数组，用标技术组更新原数组。

### 代码

```go
func setZeroes(matrix [][]int) {
  // 创建两个一维的标记数组，分别代表行和列的位置信息
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
        // 如果原数组值为0，那么分别将row和col的位置设置为true
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
        // 如果row或col对应位置值为true，那么将原数组对应位置的值设置为0
				r[j] = 0
			}
		}
	}
}
```

## 方法二：使用两个标记变量

### 思路和算法

可以用矩阵的第一行和第一列代替方法中的两个标记数组，来达到O(1)的额外空间，但是这样会导致原数组的第一行和第一列被修改，无法记录他们原本是否包含0，因此需要额外使用两个标记变量分别记录第一行和第一列原本是否包含0。

实际代码中，首先预处理两个标记变量，接着使用其他行与列来处理第一行和第一列，再然后反过来使用第一行和第一列更新其他的行与列，最后使用两个标记变量更新第一行和第一列即可。

### 代码

```go
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
  // 定义并初始化两个标记变量（第一行、列）
	row0, col0 := false, false
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
			break
		}
	}
  // 通过使用其他行与列处理第一行和第一列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
  // 通过第一行和列更新其他行和列
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
  // 如果第一行或第一列含0，更新第一行或第一列
	if row0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
```

## 方法三：使用一个标记变量

### 思路和算法

对方法二进一步优化，只使用一个标记变量记录第一列是否原本存在0。

第一列的第一个元素就可以标记第一行是否出现0。但是为了防止第一列的第一个元素被提前更新，我们需要从最后一行开始，倒序处理矩阵元素。

### 代码

```go
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	col0 := false
  // 定义并初始化标记变量，记录第一列是否原本存在
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
```

