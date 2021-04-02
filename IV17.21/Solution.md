# IV17.21直方图的水量

## 方法一：动态规划

### 思路

对于下标i，水能到达的最大高度等于下标两边的最大高度的最小值，下标i处能接的水量等于下标i处的水能到达的最大高度减去height[i]

朴素的做法是对数组height中的每个元素，分别向左右扫描并记录左边和右边的最大高度，然后计算每个下标位置能接的水的量。假设数组height的长度为n，该做法需要对每个下标位置使用O(n)的时间向两边扫描并得到最大高度，因此总时间复杂度是O(n^2)上述做法时间复杂度高是因为需要对每个下标位置都向两边扫描。如果已经知道每个位置两边的最大高度，就可以在O(n)de时间内得到能接的水的总量。使用动态规划的方法，可以在O(n)的时间内预处理得到每个位置两边的最大高度

创建两个长度为n的数组leftMax和rightMax。对于0≤i<n，leftMax[i]表示下标i及其左边的位置中height的最大高度，rightMax[i]表示下标i及其右边的位置中，height的最大高度

显然，leftMax[0]=height[0]，rightMax[n-1]=height[n-1]（左右边界）。两个数组的其余元素的计算如下：

![image-20210402092209864](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210402092209864.png)

因此可以正向遍历数组height的到数组leftMax的每个元素值，反向遍历数组height得到数组rightMax的每个元素值

得到数组leftMax和rightMax的每个元素值之后，对于0≤i<n，下标处能接到的水的量等于min（leftMax[i]，rightMax[i]）-height[i]。遍历每个下标位置就可以得到能接的水的总量

具体如图

![image-20210402092852812](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210402092852812.png)

### 代码

```go
func trap(height []int) (ans int) {
    n := len(height)
    if n == 0 {
        return
    }

    leftMax := make([]int, n)
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }

    rightMax := make([]int, n)
    rightMax[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }

    for i, h := range height {
        ans += min(leftMax[i], rightMax[i]) - h
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

