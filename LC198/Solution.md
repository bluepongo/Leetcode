# LC198 打家劫舍

## 方法一：动态规划

房屋数量一间，就直接为该间金额；两间就对比两间较大的一间。

如果房屋数量大于两间，对于第k（k>2）间房屋，有两个选项：

- 偷窃第k间房屋，就不能偷窃第k-1间房屋，偷窃总金额为前k-2间房屋的最高总金额与第k间房屋的金额之和
- 不偷窃第k间房屋，偷窃总金额就是前k-1间房屋最高总金额

在两个选项中选择偷窃总金额较大的选项，该选项对应的偷窃总金额即为前 k 间房屋能偷窃到的最高总金额。

用`dp[i]`表示前i间房屋能偷窃到的最高总金额，有如下状态转移方程：

![image-20210415102630486](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210415102630486.png)

边界条件为：

![image-20210415102645195](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210415102645195.png)

最终答案就是`dp[n-1]`其中n是数组的长度

```go
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

