# LC001 两数之和

## 方法一：暴力枚举

略

## 方法二：哈希表

### 思路及算法

注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。

使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。

这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。

### 代码

```go
func twoSum(nums []int, target int) []int {
  	// 建立哈希表并遍历
  	// key:具体值
  	// val:数字所在位置序号
    hashTable := map[int]int{}
    for i, x := range nums {
      	// 如果p
        if p, ok := hashTable[target-x]; ok {
            return []int{p, i}
        }
        hashTable[x] = i
    }
    return nil
}
```

### *Tips：哈希表循环*

```go
if _, ok := map[key]; ok {
    // 存在, _处返回key对应的val
}
```

