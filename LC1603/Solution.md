# LC1603 设计停车系统

## 方法：模拟

### 思路与算法

为每种车维护一个计数器，初始值为车位的数目。此后，每来一辆车，就将对应类型的计数器减 11。当计数器为 00 时，说明车位已满。

### 代码

```go
type ParkingSystem struct {
	left [3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{left: [3]int{big, medium, small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if this.left[carType - 1] > 0{
		this.left[carType - 1]--
		return true
	}
	return false
}
```

