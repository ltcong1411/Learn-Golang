package circle

import (
	"calculate"
	"math"
)

func DienTich(r float32) float32 {
	return float32(calculate.Multiply(calculate.Multiply(r, r), math.Pi))
}
func ChuVi(r float32) float32 {
	return float32(calculate.Multiply(calculate.Multiply(r, math.Pi), 2))
}
