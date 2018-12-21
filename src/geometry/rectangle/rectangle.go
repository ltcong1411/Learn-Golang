package rectangle

import (
	"calculate"
)

func DienTich(a, b float32) float32 {
	return calculate.Multiply(a, b)
}
func ChuVi(a, b float32) float32 {
	return calculate.Add(a, b) * 2
}
