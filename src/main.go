package main

import (
	"array/oneD"
	"fmt"
)

func main() {
	// fmt.Println(rectangle.DienTich(2, 3))
	// fmt.Println(rectangle.ChuVi(2, 4))
	// fmt.Println(circle.DienTich(2))
	// fmt.Println(circle.ChuVi(3))
	// := []int{3, 3, 6, 7, 9, 9}
	// a := []int{6, 2, -4, 6, 8, -9, 4, 5, -3, -12, 434, 54, 65, -43}
	// var t, y, u, i []int = array.Demo7(a)
	// fmt.Println(t)
	// fmt.Println(y)
	// fmt.Println(u)
	// fmt.Println(i)
	a := []int{6, 2, 4, 6, 8, 9, 4, 5, -2, -1}
	// array.Demo8()
	// fmt.Println(array.Demo7(a))
	// oneD.NguyenTo(a)
	fmt.Println(oneD.NguyenTo(a))
}
