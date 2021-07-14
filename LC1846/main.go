package main

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	length := len(arr)
	countArr := make([]int, length)
	interCount := 0 // 在计数范围内的元素个数

	// 计数(其实记录 1 到 length - 1 就够了，因为值为 length 的元素也是“万能牌”)
	for _, v := range arr {
		if v >= 1 && v <= length {
			countArr[v-1]++
			interCount++
		}
	}

	totalNeedAdd := 0 // 在遍历当前元素时，有多少比当前小的元素需要被补充
	interMax := 0     //计数范围内部的实际最大值

	//计算计数范围内部真正最大值(可优化)
	for k, v := range countArr {
		switch {
		case v == 0:
			totalNeedAdd++
		case v != 0 && v > totalNeedAdd:
			totalNeedAdd = 0
			interMax = k + 1
		case v != 0 && v <= totalNeedAdd:
			totalNeedAdd -= v
			interMax += v
		}
	}

	return interMax + (length - interCount)
}
