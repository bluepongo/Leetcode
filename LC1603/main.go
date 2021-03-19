package main

import "fmt"

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

func main(){
	parkingSystem := ParkingSystem{[3]int{1, 1, 0}}
	fmt.Println(parkingSystem.AddCar(1))
	fmt.Println(parkingSystem.AddCar(2))
	fmt.Println(parkingSystem.AddCar(3))
	fmt.Println(parkingSystem.AddCar(1))

}
