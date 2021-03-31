# LC078 子集

## 方法一：迭代法实现子集枚举

### 思路与算法

记原序列中元素总数为n，原序列中每个数字ai的状态可能是两种：「在子集中」和「不在子集中」，分别用1和0表示两种状态。每一个子集都可以对应一个长度为n的01序列，第i位表示ai是否在子集中。

| 01序列 | 子集    | 01序列对应二进制数 |
| ------ | ------- | ------------------ |
| 000    | {}      | 0                  |
| 001    | {9}     | 1                  |
| 010    | {2}     | 2                  |
| 011    | {2,9}   | 3                  |
| 100    | {5}     | 4                  |
| 101    | {5,9}   | 5                  |
| 110    | {5,2}   | 6                  |
| 111    | {5,2,9} | 7                  |

由此可以发现01序列对应的二进制数正好是从0到2^n-1。所以我们可以枚举mask∈[0,2^n-1]，mask的二进制表示是一个01序列，我们可以按照这个01序列再原集合中取数。当我们枚举完所有2^n个mask，我们也就能构造出所有的子集

### 代码

```go
func subsets(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}
```

## 方法二：递归法实现子集枚举（*）

### 思路与算法

也可以用递归来实现子集枚举

假设需要找到一个长度为n的序列a的所有子序列，代码框架如下：

```c++
vector<int> t;
void dfs(int cur, int n) {
    if (cur == n) {
        // 记录答案
        // ...
        return;
    }
    // 考虑选择当前位置
    t.push_back(cur);
    dfs(cur + 1, n, k);
    t.pop_back();
    // 考虑不选择当前位置
    dfs(cur + 1, n, k);
}
```

上面代码中，dfs(cur,n)参数表示当前位置是cur，原序列总长度为n。原序列的每个位置在答案序列的状态都是选中和不选中两种，我们用t数组存放被选中的数组。在进入dfs(cur,n)之前[0,cur-1]位置的状态是确定的，而[cur,n-1]内位置的状态是不确定的，dfs(cur,n)需要确定cur位置的状态，然后求解子问题dfs(cur+1, n)。对于cur位置，我们需要考虑a[cur]取或者不取，如果取，就把a[cur]放入一个临时的答案数组中(t)，在执行dfs(cur+1,n)，执行结束后需要对t进行回溯；如果不取，则直接执行dfs(cur+1, n)。整个递归调用的过程中，cur是从小到大递增的，当cur增加到n的时候，记录答案并终止递归。

### 代码

```go
func subsets(nums []int) (ans [][]int) {
    set := []int{}
    var dfs func(int)
    dfs = func(cur int) {
        if cur == len(nums) {
            ans = append(ans, append([]int(nil), set...))
            return
        }
        set = append(set, nums[cur])
        dfs(cur + 1)
        set = set[:len(set)-1]
        dfs(cur + 1)
    }
    dfs(0)
    return
}
```

