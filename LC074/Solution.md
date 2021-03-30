# LC074 搜索二维矩阵

## 方法一：两次二分查找

### 思路

由于每行的第一个元素大于前一行的最后一个元素，且每行元素是升序的，所以每行的第一个元素大于前一行的第一个元素，因此矩阵第一列的元素是升序的

可以对矩阵的第一列元素二分查找，找到最后一个不大于目标值的元素，然后在该元素所在行中二分查找目标值是否存在

### 代码

```go
func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
```

## 方法二：一次二分查找

### 思路

将矩阵的每一行拼接在上一行的末尾，会得到一个升序数组，我们可以再该数组上二分找到目标元素

代码实现时，可以二分升序数组的下标，将其映射到原矩阵的行和列上

### 代码

```go
func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
    return i < m*n && matrix[i/n][i%n] == target
}
```

