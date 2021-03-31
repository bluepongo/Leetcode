# LC090 子集II

## 方法一：迭代法实现子集枚举

#### 思路

考虑数组[1, 2, 2] ，选择前两个数，或者第一个、第三个数，都会得到相同的子集

也就是说，对于当前选择的数x，如果前面有与其相同的数y，且没有选择y，此时包含x的子集，必然会出现在包含y的所有子集中

我们可以通过判断这种情况，来避免生成重复的子集。先将数组排序，迭代的时候，如果发现没有选择上一个数，且当前的数字和上一个数字相同，就可以跳过当前生成的子集

### 代码

```go
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}
```

## 方法二：递归法实现子集枚举

### 思路

和方法一类似，递归的时候，如果发现没有选择上一个数，且当前数字和上一个数相同，就可以跳过当前生成的子集

### 代码

```go

func subsetsWithDup(nums []int) (ans [][]int) {
    sort.Ints(nums)
    t := []int{}
    var dfs func(bool, int)
    dfs = func(choosePre bool, cur int) {
        if cur == len(nums) {
            ans = append(ans, append([]int(nil), t...))
            return
        }
        dfs(false, cur+1)
        if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
            return
        }
        t = append(t, nums[cur])
        dfs(true, cur+1)
        t = t[:len(t)-1]
    }
    dfs(false, 0)
    return
}
```

