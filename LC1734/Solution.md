# LC1734 解码异或后的排列

## 主要思路

给定解码后数组encoded，求原始数组perm，perm是前n个正整数排列，且为奇数

因为n是奇数，所以除了perm[0]以外，数组perm还有n-1个其他元素，n-1是偶数，又因为数组encoded的每个元素都是数组perm的两个元素异或运算的结果，所以数组encoded中存在(n-1)/2个元素，这些元素的异或运算结果为数组perm除了perm[0]以外的全部元素的异或运算结果

具体来说，数组encoded的所有下标为奇数的元素的异或运算结果就是数组perm除了perm[0]以外的全部元素的异或运算结果。用odd表示数组encoded的所有下标为奇数的元素的异或运算结果

![image-20210511092124455](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210511092124455.png)

根据total和odd的值可以算出perm[0]的值

![image-20210511092147540](/Users/zhanghanyu/Library/Application Support/typora-user-images/image-20210511092147540.png)

